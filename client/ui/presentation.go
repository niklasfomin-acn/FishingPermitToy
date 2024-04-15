package presentation

import (
	"client/data"
	"image/jpeg"
	"log"
	"os"

	"github.com/rivo/tview"
)

func SetupStartPage(app *tview.Application, pages *tview.Pages) {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Willkommen im Angler Portal Berlin/Brandenburg").SetTitleAlign(tview.AlignCenter)

	file, err := os.Open("/Users/niklas.fomin/Documents/ReposLocal/FishingPermitToy/client/ui/angelschein.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	form.AddImage("", img, 75, 15, 0)

	form.AddDropDown("Nutzerstatus", []string{"Bürger", "Admin"}, 0, nil)

	// InputFields
	form.AddInputField("Nutzername", "", 30, nil, nil)
	form.AddPasswordField("Password", "", 30, '*', nil)

	// Buttons
	form.AddButton("Login", func() { pages.SwitchToPage("CitizenLandingPage") })
	form.AddButton("Beenden", func() { app.Stop() })

	form.SetButtonBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	form.SetButtonTextColor(tview.Styles.PrimaryTextColor)

	form.AddTextView("Impressum", "Accenture x BMI / LLM-supported Cloud Native Assessment / PoC Application V0.1", 50, 0, false, false).SetBorder(true).SetBorderColor(tview.Styles.BorderColor)

	pages.AddPage("StartPage", form, true, true)
}

func SetupCitizenLandingPage(app *tview.Application, pages *tview.Pages) {
	landingForm := tview.NewForm()
	landingForm.SetBorder(true).SetTitle("Menü Antragsteller").SetTitleAlign(tview.AlignCenter)
	landingForm.SetButtonBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	landingForm.SetButtonTextColor(tview.Styles.PrimaryTextColor)

	file, err := os.Open("/Users/niklas.fomin/Documents/ReposLocal/FishingPermitToy/client/ui/angelschein.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	landingForm.AddImage("", img, 70, 15, 0)

	landingForm.AddButton("Ausweisdokument Hochladen", func() { TriggerDocumentAI() })
	landingForm.AddButton("Manuellen Antrag Erstellen", func() { pages.SwitchToPage("ManualPermitPage") })
	landingForm.AddButton("Zurück", func() { pages.SwitchToPage("StartPage") })
	landingForm.AddButton("Beenden", func() { app.Stop() })
	pages.AddPage("CitizenLandingPage", landingForm, true, false)
	landingForm.AddTextView("Impressum", "Accenture x BMI / LLM-supported Cloud Native Assessment / PoC Application V0.1", 50, 0, false, false).SetBorder(true).SetBorderColor(tview.Styles.BorderColor)
}

func SetupManualPermitPage(app *tview.Application, pages *tview.Pages) {
	manualPermitForm := tview.NewForm()
	manualPermitForm.SetBorder(true).SetTitle("Manueller Antrag").SetTitleAlign(tview.AlignCenter)

	file, err := os.Open("/Users/niklas.fomin/Documents/ReposLocal/FishingPermitToy/client/ui/perso.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	manualPermitForm.AddImage("", img, 40, 10, 0)

	manualPermitForm.AddInputField("Personalausweis-Nr.", "", 30, nil, nil)
	manualPermitForm.AddInputField("Nachname", "", 30, nil, nil)
	manualPermitForm.AddInputField("Vorname", "", 30, nil, nil)
	manualPermitForm.AddInputField("Geburtsdatum [dd:mm:yy]", "", 30, nil, nil)
	manualPermitForm.AddInputField("Geburtsort", "", 30, nil, nil)
	manualPermitForm.AddInputField("Geschlecht", "", 30, nil, nil)
	manualPermitForm.AddInputField("Nationalität", "", 30, nil, nil)
	manualPermitForm.AddInputField("Datum der Ausstellung [dd:mm:yy]", "", 30, nil, nil)
	manualPermitForm.AddInputField("Gültig bis: [dd:mm:yy]", "", 30, nil, nil)
	manualPermitForm.AddInputField("Ausstellende Behörde [Amt]", "", 30, nil, nil)
	manualPermitForm.AddInputField("Antragstyp[Hobby/Gewerblich]", "", 30, nil, nil)
	manualPermitForm.AddInputField("Vorgesehene Nutzungsregion [Berlin/Brandenburg]", "", 30, nil, nil)
	manualPermitForm.AddInputField("E-Mail Adresse [max.mustermann@xxx.com]", "", 30, nil, nil)
	manualPermitForm.AddInputField("Telefonnummer[+49 12345678]", "", 30, nil, nil)

	manualPermitForm.AddButton("Manuellen Neuen Antrag stellen", func() {
		citizenPermit := data.CreateCitizenPermitFromForm(manualPermitForm)
		// Call API handler to post citizen permit request
		requestClient := data.NewJSONTransferClient("localhost", "3000", "/citizenPermit")
		requestClient.TransferCitizenPermit(citizenPermit)
	})

	pages.AddPage("ManualPermitPage", manualPermitForm, true, false)

	manualPermitForm.AddButton("Zurück", func() { pages.SwitchToPage("CitizenLandingPage") })
	manualPermitForm.AddButton("Beenden", func() { app.Stop() })
}

func SetupAdminPage(app *tview.Application, pages *tview.Pages) {
	// Buttons
	adminForm := tview.NewForm()
	adminForm.SetBorder(true).SetTitle("Admin Menue").SetTitleAlign(tview.AlignCenter)
	adminForm.AddButton("Zurück", func() { pages.SwitchToPage("StartPage") })
	adminForm.AddButton("Beenden", func() { app.Stop() })
	adminForm.AddButton("Anträge anzeigen", func() { pages.SwitchToPage("PermitPage") })

	pages.AddPage("AdminPage", adminForm, true, false)

}

func SetupPermitPage(app *tview.Application, pages *tview.Pages) {
	// Buttons
	permitForm := tview.NewForm()
	permitForm.SetBorder(true).SetTitle("Permit Menue").SetTitleAlign(tview.AlignCenter)
	permitForm.AddButton("Zurück", func() { pages.SwitchToPage("AdminPage") })
	permitForm.AddButton("Beenden", func() { app.Stop() })
	permitForm.AddButton("Anträge anzeigen", func() {})
	permitForm.AddButton("Antrag genehmigen", func() {})
	permitForm.AddButton("Antrag ablehnen", func() {})

	pages.AddPage("PermitPage", permitForm, true, false)

}

func TriggerDocumentAI() {}
