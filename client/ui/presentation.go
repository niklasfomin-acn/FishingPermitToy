package presentation

import (
	"fmt"
	"log"

	"client/data"

	"github.com/mbndr/figlet4go"
	"github.com/rivo/tview"
)

// func InitClient() {
// 	app := tview.NewApplication()
// }

func ShowWelcomeMessage() {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontName = "lean"

	result, err := ascii.Render("Angelschein Portal")
	if err != nil {
		log.Fatalf("Error rendering ASCII: %v", err)
	}
	fmt.Println(result)
}

// func WelcomePage() {
// 	// Design with Logo and branding
// 	// Time and Date
// 	// Login Mask
// 	//
// }

func ShowOptions() {

	//TODO: Change Color Schema
	app := tview.NewApplication()
	form := tview.NewForm()

	// Input Fields
	form.AddInputField("Passport Number", "", 30, nil, nil)
	form.AddInputField("Surname", "", 30, nil, nil)
	form.AddInputField("Given Names", "", 30, nil, nil)
	form.AddInputField("Date of Birth", "", 30, nil, nil)
	form.AddInputField("Place of Birth", "", 30, nil, nil)
	form.AddInputField("Gender", "", 30, nil, nil)
	form.AddInputField("Nationality", "", 30, nil, nil)
	form.AddInputField("Date of Issue", "", 30, nil, nil)
	form.AddInputField("Expiry Date", "", 30, nil, nil)
	form.AddInputField("Issuing Authority", "", 30, nil, nil)
	form.AddInputField("Permit Type", "", 30, nil, nil)
	form.AddInputField("Permit Location", "", 30, nil, nil)

	// Buttons
	form.AddButton("Neuen Antrag stellen", func() {
		citizenPermit := data.CreateCitizenPermitFromForm(form)
		// Call API handler to post citizen permit request
		requestClient := data.NewJSONTransferClient("localhost", "3000", "/citizenPermit")
		requestClient.TransferCitizenPermit(citizenPermit)
	})
	form.AddButton("Antragstatus abfragen", func() {
		// Angaben zum Antragsteller
		citizenPermit := data.CreateCitizenPermitFromForm(form)
		citizenPermitView := tview.NewTextView().SetText(fmt.Sprintf("%+v", citizenPermit))

		flex := tview.NewFlex().SetDirection(tview.FlexRow)
		flex.AddItem(citizenPermitView, 0, 1, false)
		app.SetRoot(flex, true)

	})

	// TODO: Restructure the whole app view. Only admin can call this.
	form.AddButton("View Permits", func() {
		adminClient, err := data.NewGrpcClient(":50051")
		if err != nil {
			log.Fatalf("Error creating new grpc client: %v", err)
		}
		citizenPermits, err := adminClient.FetchCitizenPermits()
		if err != nil {
			log.Fatalf("Error fetching citizen permits: %v", err)
		}

		permitText := ""
		for _, permit := range citizenPermits {
			permitText += fmt.Sprintf("%+v\n", permit)
		}

		permitView := tview.NewTextView().SetText(permitText)

		flex := tview.NewFlex().SetDirection(tview.FlexRow)
		flex.AddItem(permitView, 0, 1, false)
		app.SetRoot(flex, true)

	})
	form.AddButton("Beenden", func() {
		app.Stop()
	})

	// Layout
	form.SetBorder(true).SetTitle("Willkommen beim Angelschein Portal Berlin!").SetTitleAlign(tview.AlignLeft)

	if err := app.SetRoot(form, true).Run(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}

}
