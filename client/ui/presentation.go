package presentation

import (
	"fmt"
	"log"

	"github.com/mbndr/figlet4go"
	"github.com/rivo/tview"

	"client/data"
)

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

func ShowOptions() {

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
		//citizen := data.CreateCitizenFromForm(form)
		permit := data.CreatePermitFromForm(form)

		// Call API handler to post permit request
		client := data.NewJSONTransferClient("localhost", "8080", "/permit")
		client.TransferPermit(permit)

	})
	form.AddButton("Antragstatus abfragen", func() {
		// Angaben zum Antrag
		permit := data.CreatePermitFromForm(form)
		permitView := tview.NewTextView().SetText(fmt.Sprintf("%+v", permit))
		// Angaben zum Antragsteller
		citizen := data.CreateCitizenFromForm(form)
		citizenView := tview.NewTextView().SetText(fmt.Sprintf("%+v", citizen))

		flex := tview.NewFlex().SetDirection(tview.FlexRow)
		flex.AddItem(permitView, 0, 1, false)
		flex.AddItem(citizenView, 0, 1, false)
		app.SetRoot(flex, true)

	})

	form.AddButton("Antrag zurückziehen", func() {
		app.Stop()
	})
	form.AddButton("Beenden", func() {
		app.Stop()
	})

	// Layout
	form.SetBorder(true).SetTitle("Willkommen beim Angelschein Portal Berlin!").SetTitleAlign(tview.AlignLeft)

	if err := app.SetRoot(form, true).Run(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}

	//currentTimeView := tview.NewTextView()

	// Doesn't work
	// go func() {
	// 	for {
	// 		currenttime := time.Now().Format("15:04:05")
	// 		app.QueueUpdateDraw(func() {
	// 			currentTimeView.SetText(currenttime)
	// 		})
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	// form.AddFormItem(currentTimeView)
}
