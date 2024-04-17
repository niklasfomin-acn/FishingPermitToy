package presentation

import (
	"image/jpeg"
	"log"
	"os"

	"client/data"
	config "client/data"

	"github.com/rivo/tview"
)

func SetupStartPage(app *tview.Application, pages *tview.Pages) {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Willkommen im Angelschein-Portal Berlin-Brandenburg").SetTitleAlign(tview.AlignCenter)
	form.SetButtonBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	form.SetButtonTextColor(tview.Styles.PrimaryTextColor)

	// Header
	file, err := os.Open("/Users/niklas.fomin/Documents/ReposLocal/FishingPermitToy/client/ui/angelschein.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	form.AddImage("", img, 75, 15, 0)

	// InputFields
	form.AddDropDown("Nutzerstatus", []string{"Bürger", "Administrator"}, 0, nil)
	form.AddInputField("Nutzername", "", 30, nil, nil)
	form.AddPasswordField("Password", "", 30, '*', nil)

	// Buttons
	form.AddButton("Login", func() {
		if _, option := form.GetFormItemByLabel("Nutzerstatus").(*tview.DropDown).GetCurrentOption(); option == "Bürger" {
			pages.SwitchToPage("CitizenLandingPage")
		} else {
			pages.SwitchToPage("AdminPage")
		}
	})

	form.AddButton("Beenden", func() { app.Stop() })

	// Footer
	form.AddTextView("Impressum", " > Accenture Technology - all rights reserved. / Accenture x BMI / Project: LLM-supported Cloud Native Assessment / MIT-licensed PoC Application V0.3 / StaffDevs: Purwins, Hendrick, Gruhl, Constantin, Fomin, Niklas", 50, 0, false, false).SetBorder(true).SetBorderColor(tview.Styles.BorderColor)

	pages.AddPage("StartPage", form, true, true)
}

func SetupCitizenLandingPage(app *tview.Application, pages *tview.Pages, config config.Config) {
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

	// Header
	landingForm.AddImage("", img, 70, 15, 0)

	// Buttons
	landingForm.AddButton("Antragstatus Anzeigen", func() {})
	landingForm.AddButton("Ausweisdokument Hochladen", func() { TriggerDocumentAI() })
	landingForm.AddButton("Manuellen Antrag Erstellen", func() { pages.SwitchToPage("ManualPermitPage") })
	landingForm.AddButton("Zurück", func() { pages.SwitchToPage("StartPage") })
	landingForm.AddButton("Beenden", func() { app.Stop() })

	pages.AddPage("CitizenLandingPage", landingForm, true, false)
	landingForm.AddTextView("Impressum", "Accenture x BMI / LLM-supported Cloud Native Assessment / PoC Application V0.3", 50, 0, false, false).SetBorder(true).SetBorderColor(tview.Styles.BorderColor)
}

func SetupManualPermitPage(app *tview.Application, pages *tview.Pages, config config.Config) {
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
		requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPI)
		requestClient.TransferCitizenPermit(citizenPermit)
	})

	pages.AddPage("ManualPermitPage", manualPermitForm, true, false)

	manualPermitForm.AddButton("Zurück", func() { pages.SwitchToPage("CitizenLandingPage") })
	manualPermitForm.AddButton("Beenden", func() { app.Stop() })
}

func SetupAdminPage(app *tview.Application, pages *tview.Pages, config config.Config) {
	// Layout
	adminForm := tview.NewForm()
	adminForm.SetBorder(true).SetTitle("Admin Menue").SetTitleAlign(tview.AlignCenter)

	// Buttons
	adminForm.AddButton("Anträge anzeigen", func() { pages.SwitchToPage("PermitPage") })
	adminForm.AddButton("Zurück", func() { pages.SwitchToPage("StartPage") })
	adminForm.AddButton("Beenden", func() { app.Stop() })

	pages.AddPage("AdminPage", adminForm, true, false)

}

func SetupPermitPage(app *tview.Application, pages *tview.Pages, config config.Config) {
	// Layout
	permitForm := tview.NewForm()
	permitForm.SetBorder(true).SetTitle("Anträge bearbeiten").SetTitleAlign(tview.AlignCenter)

	// Buttons
	permitForm.AddButton("Alle Anträge anzeigen", func() { pages.SwitchToPage("AdminPage") })
	permitForm.AddButton("Zurück", func() { pages.SwitchToPage("AdminPage") })
	permitForm.AddButton("Beenden", func() { app.Stop() })

	pages.AddPage("PermitPage", permitForm, true, false)

}

func TriggerDocumentAI() {}
