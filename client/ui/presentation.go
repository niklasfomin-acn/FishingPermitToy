package presentation

import (
	"client/data"
	config "client/data"
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Login
func SetupStartPage(app *tview.Application, pages *tview.Pages) {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Willkommen im Angelschein-Portal Berlin-Brandenburg").SetTitleAlign(tview.AlignCenter)
	form.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	// Header
	// file, err := os.Open("/Users/niklas.fomin/Documents/ReposLocal/FishingPermitToy/client/ui/angelschein.jpeg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// img, err := jpeg.Decode(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// form.AddImage("", img, 75, 15, 0)

	// InputFields
	form.AddDropDown("Nutzerstatus", []string{"Bürger", "Administrator"}, 0, nil)
	form.AddInputField("Nutzername", "", 30, nil, nil)
	form.AddPasswordField("Password", "", 30, '*', nil)

	// Buttons
	form.AddButton("Login", func() {
		if _, option := form.GetFormItemByLabel("Nutzerstatus").(*tview.DropDown).GetCurrentOption(); option == "Bürger" {
			//TODO: Implement keycloak auhtentication for role citizen
			pages.SwitchToPage("CitizenLandingPage")
		} else {
			//TODO: Implement keycloak auhtentication for role admin
			pages.SwitchToPage("AdminPage")
		}
	})

	form.AddButton("Logout", func() {
		//TODO: Implement keycloak logout
	})
	form.AddButton("Beenden", func() { app.Stop() })

	// Footer
	form.AddTextView("Impressum", " > Accenture Technology - all rights reserved. Accenture x BMI / Project: LLM-supported Cloud Native Assessment / MIT-licensed PoC Application V0.3 / StaffDevs: Fomin, Niklas", 50, 0, false, false).SetBorder(true).SetBorderColor(tview.Styles.BorderColor)

	// Finish Page Setup
	pages.AddPage("StartPage", form, true, true)
}

// Citizen Options
func SetupCitizenLandingPage(app *tview.Application, pages *tview.Pages, config config.Config) {
	landingForm := tview.NewForm()
	landingForm.SetBorder(true).SetTitle("Menü Antragsteller").SetTitleAlign(tview.AlignCenter)
	landingForm.SetButtonTextColor(tview.Styles.PrimaryTextColor)

	// file, err := os.Open("/Users/niklas.fomin/Documents/ReposLocal/FishingPermitToy/client/ui/angelschein.jpeg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// img, err := jpeg.Decode(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Header
	// landingForm.AddImage("", img, 70, 15, 0)

	// Buttons
	landingForm.AddButton("Antragstatus Anzeigen", func() {})
	landingForm.AddButton("Ausweisdokument Hochladen", func() {})
	landingForm.AddButton("Manuellen Antrag Erstellen", func() { pages.SwitchToPage("ManualPermitPage") })
	landingForm.AddButton("Zurück", func() { pages.SwitchToPage("StartPage") })
	landingForm.AddButton("Beenden", func() { app.Stop() })

	// Footer
	landingForm.AddTextView("Impressum", "Accenture x BMI / LLM-supported Cloud Native Assessment / PoC Application V0.3", 50, 0, false, false).SetBorder(true).SetBorderColor(tview.Styles.BorderColor)

	// Finish Page Setup
	pages.AddPage("CitizenLandingPage", landingForm, true, false)

}

// Citizen Options for Manual Permit Creation
func SetupManualPermitPage(app *tview.Application, pages *tview.Pages, config config.Config) {
	manualPermitForm := tview.NewForm()
	manualPermitForm.SetBorder(true).SetTitle("Manueller Antrag").SetTitleAlign(tview.AlignCenter)

	// file, err := os.Open("/Users/niklas.fomin/Documents/ReposLocal/FishingPermitToy/client/ui/perso.jpeg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// img, err := jpeg.Decode(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// manualPermitForm.AddImage("", img, 40, 10, 0)

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
		requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[0])
		requestClient.TransferCitizenPermit(citizenPermit)
	})

	// Buttons
	manualPermitForm.AddButton("Zurück", func() { pages.SwitchToPage("CitizenLandingPage") })
	manualPermitForm.AddButton("Beenden", func() { app.Stop() })

	// Finish Page Setup
	pages.AddPage("ManualPermitPage", manualPermitForm, true, false)
}

// Admin Options
func SetupAdminPage(app *tview.Application, pages *tview.Pages, config config.Config) {
	// Layout
	adminForm := tview.NewForm()
	adminForm.SetBorder(true).SetTitle("Administrator Menü").SetTitleAlign(tview.AlignCenter)
	adminForm.AddTextView("Impressum", "Accenture x BMI / LLM-supported Cloud Native Assessment / PoC Application V0.3", 100, 0, false, false).SetBorder(true).SetBorderColor(tview.Styles.BorderColor)

	// Buttons
	adminForm.AddButton("Alle Anträge anzeigen", func() {
		// Call API handler to fetch all citizen permit requests
		requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[1])
		allCitizenPermitRequests, err := requestClient.ViewAllCitizenPermitRequests()
		if err != nil {
			log.Printf("Error fetching all citizen permit requests: %v", err)
		} else {
			// Create a List Item
			list := tview.NewList()
			details := tview.NewTextView()
			details.SetTextColor(tcell.ColorYellow)
			details.SetBorder(true).SetTitle("Informationen zum Antrag").
				SetTitleAlign(tview.AlignCenter).SetBackgroundColor(tcell.ColorDefault)

			list.AddItem("Zurück", "", 'b', func() {
				pages.SwitchToPage("AdminPage")
				details.SetText("")
			})
			// Write the fetched data to the list
			listCounter := 1
			for _, cp := range allCitizenPermitRequests {
				secondaryText := fmt.Sprintf("%v\n								Personalausweis-Nummer: %v\n								Antrag-Status: %v\n								Geburtsdatum: %v\n								Antrags-Typ: %v\n								Ort der Nutzung: %v\n								E-Mail: %v\n								Telefonnummer: %v\n								Ausstellungsdatum: %v\n								Ablaufdatum: %v\n								Behörde: %v\n								Staatsangehörigkeit: %v", "", cp.PassportNumber, cp.PermitState, cp.DateOfBirth, cp.PermitType, cp.PermitLocation, cp.Email, cp.Phone, cp.DateOfIssue, cp.ExpiryDate, cp.IssuingAuthority, cp.Nationality)
				//list.AddItem(fmt.Sprintf("%d - %s", listCounter, cp.Surname), secondaryText, 0, nil)
				list.AddItem(fmt.Sprintf("%d - %s", listCounter, cp.Surname), "", 0, func() {
					details.SetText(secondaryText)
				})
				listCounter++
			}

			flex := tview.NewFlex().AddItem(list, 0, 1, true).AddItem(details, 0, 2, false)
			pages.AddPage("CitizenPermitview", flex, true, true)
			pages.SwitchToPage("CitizenPermitview")
		}
	})
	adminForm.AddButton("Zu Bearbeiten", func() {
		// Call API handler to fetch all citizen permit requests
		requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[2])
		allCitizenPermitRequests, err := requestClient.ViewPendingCitizenPermitRequests()
		if err != nil {
			log.Printf("Error fetching pending citizen permit requests: %v", err)
		} else {
			// Create a List Item
			list := tview.NewList()
			details := tview.NewTextView()
			details.SetTextColor(tcell.ColorYellow)
			list.AddItem("Zurück", "", 'b', func() {
				pages.SwitchToPage("AdminPage")
			})
			// Write the fetched data to the list
			listCounter := 1
			for _, cp := range allCitizenPermitRequests {
				secondaryText := fmt.Sprintf("Personalausweis-Nummer: %v\nAntrag-Status: %v\nGeburtsdatum: %v\nAntrags-Typ: %v\nOrt der Nutzung: %v\nE-Mail: %v\nTelefonnummer: %v\nAusstellungsdatum: %v\nAblaufdatum: %v\nBehörde: %v\nStaatsangehörigkeit: %v\n", cp.PassportNumber, cp.PermitState, cp.DateOfBirth, cp.PermitType, cp.PermitLocation, cp.Email, cp.Phone, cp.DateOfIssue, cp.ExpiryDate, cp.IssuingAuthority, cp.Nationality)
				list.AddItem(fmt.Sprintf("%d - %s", listCounter, cp.Surname), "", 0, func() {
					detailsOnNewPage := tview.NewTextView().SetText(secondaryText).SetTextColor(tcell.ColorYellow)
					admissionForm := tview.NewForm()
					admissionForm.AddButton("Antrag Genehmigen", func() {
						passportIdentifier := cp.PassportNumber
						requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[6])
						requestClient.ApproveCitizenPermitRequest(passportIdentifier)

						// Notify the user
						flex := tview.NewFlex()
						notifier := tview.NewForm()
						notifier.AddButton("Zurück", func() { pages.SwitchToPage("AdminPage") })
						NotifyView := tview.NewTextView()
						NotifyView.SetText("Antrag wurde genehmigt. E-Mail Benachrichtigung wurde versendet.").SetTextColor(tcell.ColorGreen)
						flex.AddItem(NotifyView, 0, 1, false)
						flex.AddItem(notifier, 0, 2, true)
						pages.AddAndSwitchToPage("NotifyPage", flex, true)
					})
					admissionForm.AddButton("Antrag Ablehnen", func() {
						passportIdentifier := cp.PassportNumber
						requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[7])
						requestClient.RejectCitizenPermitRequest(passportIdentifier)
					})
					admissionForm.AddButton("Zurück", func() { pages.SwitchToPage("PendingCitizenPermitview") })

					flex := tview.NewFlex().SetDirection(tview.FlexRow)
					flex.AddItem(detailsOnNewPage, 0, 1, false)
					flex.AddItem(admissionForm, 0, 1, true)

					pages.AddAndSwitchToPage("ApproveOrRejectPage", flex, true)
				})
				listCounter++
			}

			flex := tview.NewFlex().AddItem(list, 0, 1, true).AddItem(details, 0, 2, false)
			pages.AddPage("PendingCitizenPermitview", flex, true, true)
			pages.SwitchToPage("PendingCitizenPermitview")
		}
	})
	adminForm.AddButton("Bearbeitete Anträge", func() {
		// Call API handler to fetch all citizen permit requests
		requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[3])
		allCitizenPermitRequests, err := requestClient.ViewProcessedCitizenPermitRequests()
		if err != nil {
			log.Printf("Error fetching processed citizen permit requests: %v", err)
		} else {
			// Create a List Item
			list := tview.NewList()
			details := tview.NewTextView()
			details.SetTextColor(tcell.ColorYellow)
			details.SetBorder(true).SetTitle("Informationen zum Antrag").
				SetTitleAlign(tview.AlignCenter).SetBackgroundColor(tcell.ColorDefault)

			list.AddItem("Zurück", "", 'b', func() {
				pages.SwitchToPage("AdminPage")
				details.SetText("")
			})
			// Write the fetched data to the list
			listCounter := 1
			for _, cp := range allCitizenPermitRequests {
				secondaryText := fmt.Sprintf("%v\n								Personalausweis-Nummer: %v\n								Antrag-Status: %v\n								Geburtsdatum: %v\n								Antrags-Typ: %v\n								Ort der Nutzung: %v\n								E-Mail: %v\n								Telefonnummer: %v\n								Ausstellungsdatum: %v\n								Ablaufdatum: %v\n								Behörde: %v\n								Staatsangehörigkeit: %v", "", cp.PassportNumber, cp.PermitState, cp.DateOfBirth, cp.PermitType, cp.PermitLocation, cp.Email, cp.Phone, cp.DateOfIssue, cp.ExpiryDate, cp.IssuingAuthority, cp.Nationality)
				//list.AddItem(fmt.Sprintf("%d - %s", listCounter, cp.Surname), secondaryText, 0, nil)
				list.AddItem(fmt.Sprintf("%d - %s", listCounter, cp.Surname), "", 0, func() {
					details.SetText(secondaryText)
				})
				listCounter++
			}

			flex := tview.NewFlex().AddItem(list, 0, 1, true).AddItem(details, 0, 2, false)
			pages.AddPage("ProcessedCitizenPermitview", flex, true, true)
			pages.SwitchToPage("ProcessedCitizenPermitview")
		}
	})
	adminForm.AddButton("Genehmigte Anträge", func() {
		// Call API handler to fetch all citizen permit requests
		requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[4])
		allCitizenPermitRequests, err := requestClient.ViewApprovedCitizenPermitRequests()
		if err != nil {
			log.Printf("Error fetching processed citizen permit requests: %v", err)
		} else {
			// Create a List Item
			list := tview.NewList()
			details := tview.NewTextView()
			details.SetTextColor(tcell.ColorYellow)
			details.SetBorder(true).SetTitle("Informationen zum Antrag").
				SetTitleAlign(tview.AlignCenter).SetBackgroundColor(tcell.ColorDefault)

			list.AddItem("Zurück", "", 'b', func() {
				pages.SwitchToPage("AdminPage")
				details.SetText("")
			})
			// Write the fetched data to the list
			listCounter := 1
			for _, cp := range allCitizenPermitRequests {
				secondaryText := fmt.Sprintf("%v\n								Personalausweis-Nummer: %v\n								Antrag-Status: %v\n								Geburtsdatum: %v\n								Antrags-Typ: %v\n								Ort der Nutzung: %v\n								E-Mail: %v\n								Telefonnummer: %v\n								Ausstellungsdatum: %v\n								Ablaufdatum: %v\n								Behörde: %v\n								Staatsangehörigkeit: %v", "", cp.PassportNumber, cp.PermitState, cp.DateOfBirth, cp.PermitType, cp.PermitLocation, cp.Email, cp.Phone, cp.DateOfIssue, cp.ExpiryDate, cp.IssuingAuthority, cp.Nationality)
				//list.AddItem(fmt.Sprintf("%d - %s", listCounter, cp.Surname), secondaryText, 0, nil)
				list.AddItem(fmt.Sprintf("%d - %s", listCounter, cp.Surname), "", 0, func() {
					details.SetText(secondaryText)
				})
				listCounter++
			}

			flex := tview.NewFlex().AddItem(list, 0, 1, true).AddItem(details, 0, 2, false)
			pages.AddPage("ApprovedCitizenPermitview", flex, true, true)
			pages.SwitchToPage("ApprovedCitizenPermitview")
		}
	})
	adminForm.AddButton("Abgelehnte Anträge", func() {
		// Call API handler to fetch all citizen permit requests
		requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[5])
		allCitizenPermitRequests, err := requestClient.ViewRejectedCitizenPermitRequests()
		if err != nil {
			log.Printf("Error fetching processed citizen permit requests: %v", err)
		} else {
			// Create a List Item
			list := tview.NewList()
			details := tview.NewTextView()
			details.SetTextColor(tcell.ColorYellow)
			details.SetBorder(true).SetTitle("Informationen zum Antrag").
				SetTitleAlign(tview.AlignCenter).SetBackgroundColor(tcell.ColorDefault)

			list.AddItem("Zurück", "", 'b', func() {
				pages.SwitchToPage("AdminPage")
				details.SetText("")
			})
			// Write the fetched data to the list
			listCounter := 1
			for _, cp := range allCitizenPermitRequests {
				secondaryText := fmt.Sprintf("%v\n								Personalausweis-Nummer: %v\n								Antrag-Status: %v\n								Geburtsdatum: %v\n								Antrags-Typ: %v\n								Ort der Nutzung: %v\n								E-Mail: %v\n								Telefonnummer: %v\n								Ausstellungsdatum: %v\n								Ablaufdatum: %v\n								Behörde: %v\n								Staatsangehörigkeit: %v", "", cp.PassportNumber, cp.PermitState, cp.DateOfBirth, cp.PermitType, cp.PermitLocation, cp.Email, cp.Phone, cp.DateOfIssue, cp.ExpiryDate, cp.IssuingAuthority, cp.Nationality)
				//list.AddItem(fmt.Sprintf("%d - %s", listCounter, cp.Surname), secondaryText, 0, nil)
				list.AddItem(fmt.Sprintf("%d - %s", listCounter, cp.Surname), "", 0, func() {
					details.SetText(secondaryText)
				})
				listCounter++
			}

			flex := tview.NewFlex().AddItem(list, 0, 1, true).AddItem(details, 0, 2, false)
			pages.AddPage("RejectedCitizenPermitview", flex, true, true)
			pages.SwitchToPage("RejectedCitizenPermitview")
		}
	})
	adminForm.AddButton("Zurück", func() { pages.SwitchToPage("StartPage") })

	// Finsih Page Setup
	pages.AddPage("AdminPage", adminForm, true, false)

}
