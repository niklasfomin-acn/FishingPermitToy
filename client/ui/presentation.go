package presentation

import (
	"client/auth"
	"client/data"
	config "client/data"
	"client/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Global Variables
var ProcessedResultMap map[string]interface{}
var Vendor string
var APIAccessToken string

// Login
func SetupStartPage(app *tview.Application, pages *tview.Pages, config config.Config) {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Willkommen im Angelschein-Portal Berlin-Brandenburg").SetTitleAlign(tview.AlignCenter)
	form.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	// InputFields
	form.AddDropDown("Nutzerstatus", []string{"Bürger", "Administrator"}, 0, nil)
	form.AddInputField("Nutzername", "", 30, nil, nil)
	form.AddPasswordField("Password", "", 30, '*', nil)

	// Buttons
	form.AddButton("Login", func() {
		// Logger for Debugging
		logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening log file: %v", err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)

		if _, option := form.GetFormItemByLabel("Nutzerstatus").(*tview.DropDown).GetCurrentOption(); option == "Bürger" {

			clientID, endpoint, secret, _, err := auth.GetIdentityFromVault(config)
			if err != nil {
				log.Fatalf("Error getting identity from vault: %v", err)
			}
			log.Printf("Keycloak Client: %v, %v, %v", clientID, secret, endpoint)

			keycloak := auth.NewKeycloakClient(clientID, endpoint, secret)
			log.Printf("Keycloak Client: %v", keycloak)

			username, password := auth.GetCredentialsFromForm(form)
			//pages.SwitchToPage("CitizenLandingPage")

			log.Printf("User Credentials: %v %v", username, password)

			_, _, _, realm, err := auth.GetIdentityFromVault(config)
			if err != nil {
				log.Fatalf("Error getting identity from vault: %v", err)
			}

			token, err := keycloak.GetToken(username, password, realm)
			if err != nil {
				log.Printf("Error getting token: %v", err)
			}

			if token {
				//APIAccessToken = token
				pages.SwitchToPage("CitizenLandingPage")
			} else {
				log.Fatalf("User Credentials are wrong: %v", err)
				pages.SwitchToPage("StartPage")
			}

		} else {
			clientID, secret, endpoint, _, err := auth.GetIdentityFromVault(config)
			if err != nil {
				log.Fatalf("Error getting identity from vault: %v", err)
			}

			keycloak := auth.NewKeycloakClient(clientID, secret, endpoint)
			log.Printf("Keycloak Client: %v", keycloak)

			username, password := auth.GetCredentialsFromForm(form)
			//pages.SwitchToPage("CitizenLandingPage")

			log.Printf("User Credentials: %v %v", username, password)

			_, _, _, realm, err := auth.GetIdentityFromVault(config)
			if err != nil {
				log.Fatalf("Error getting identity from vault: %v", err)
			}

			token, err := keycloak.GetToken(username, password, realm)
			if err != nil {
				log.Printf("Error getting token: %v", err)
			}

			if token {
				//APIAccessToken = token
				pages.SwitchToPage("AdminPage")
			} else {
				log.Fatalf("User Credentials are wrong: %v", err)
				pages.SwitchToPage("StartPage")
			}
		}
	})

	form.AddButton("Registrieren", func() {})
	form.AddButton("Logout", func() {
		//TODO: Implement keycloak logout
	})
	form.AddButton("Beenden", func() { app.Stop() })

	// Footer
	form.AddTextView("Impressum", " > Accenture Technology - all rights reserved. Accenture x BMI / Project: LLM-supported Cloud Native Assessment / MIT-licensed PoC Application V0.3 / StaffDev: Fomin, Niklas", 50, 0, false, false).SetBorder(true).SetBorderColor(tview.Styles.BorderColor)

	// Finish Page Setup
	pages.AddPage("StartPage", form, true, true)
}

// Citizen Options
func SetupCitizenLandingPage(app *tview.Application, pages *tview.Pages, config config.Config) {
	landingForm := tview.NewForm()
	landingForm.SetBorder(true).SetTitle("Menü Antragsteller").SetTitleAlign(tview.AlignCenter)
	landingForm.SetButtonTextColor(tview.Styles.PrimaryTextColor)

	// Buttons
	landingForm.AddButton("Antragstatus Anzeigen", func() {})
	landingForm.AddButton("Smarte Antragstellung", func() { pages.SwitchToPage("SmartDocumentPage") })
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
	// Logger for Debugging
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Layout
	manualPermitForm := tview.NewForm()
	manualPermitForm.SetBorder(true).SetTitle("Manueller Antrag").SetTitleAlign(tview.AlignCenter)
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

	manualPermitForm.AddButton("Neuen Antrag stellen", func() {
		citizenPermit := data.CreateCitizenPermitFromForm(manualPermitForm)
		log.Printf("Successfully created citizen permit: %v", citizenPermit)

		// Call API handler to post citizen permit request
		requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[0])
		log.Printf("Successfully created request client: %v", requestClient)
		requestClient.TransferCitizenPermit(citizenPermit, APIAccessToken)
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

// Smart Document Processing through Azure and AWS services
func SetupSmartDocumentPage(app *tview.Application, pages *tview.Pages, config config.Config) {
	smartDocumentForm := tview.NewForm()
	smartDocumentForm.SetBorder(true).SetTitle("Smarte Dokumentenverarbeitung").SetTitleAlign(tview.AlignCenter)

	// Persodokumente hardcoded TODO: Dynamische File Auswahl in der main implementieren
	var personalausweis1 string = config.FilePaths[0]
	var personalausweis2 string = config.FilePaths[1]

	// Buttons
	smartDocumentForm.AddDropDown("Vendor wählen", []string{"Azure", "AWS"}, 0, nil)
	smartDocumentForm.AddDropDown("Dokument Wählen", []string{personalausweis1, personalausweis2}, 0, nil)
	smartDocumentForm.AddButton("Dokument Hochladen", func() {

		if _, option := smartDocumentForm.GetFormItemByLabel("Vendor wählen").(*tview.DropDown).GetCurrentOption(); option == "Azure" {
			// Show waiting notification on the screen
			modal := tview.NewModal().
				SetText("Dokument bereit zur Verarbeitung...!").
				AddButtons([]string{"Mit ENTER bestätigen"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					// Define the page that occurs when clicking the button
					resultsView := tview.NewTextView()
					resultsView.SetBorder(true).SetTitle("Ergebnisse der Dokumentenverarbeitung").SetTitleAlign(tview.AlignCenter)
					//resultsView.SetText("Bitte prüfen Sie die Ergebnisse der Dokumentenverarbeitung auf Korrektheit.")
					resultsView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
						if event.Key() == tcell.KeyEnter {
							pages.SwitchToPage("SmartPermitPage")
						}
						return event
					})
					// Logger for Debugging

					logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
					if err != nil {
						log.Fatalf("Error opening log file: %v", err)
					}
					defer logFile.Close()
					log.SetOutput(logFile)

					// Process the document through AI service
					_, indexValue := smartDocumentForm.GetFormItemByLabel("Dokument Wählen").(*tview.DropDown).GetCurrentOption()
					// Access the keyvault here, TODO: Exchange bools versus config.json stuff
					pathFromVault, secretFromVault, err := auth.GetSecretFromVault(true, false, false, &config)
					// if err != nil {
					// 	log.Fatalf("Error getting secret from vault: %v", err)
					// }
					// log.Printf("Secret from Vault: %v", secretFromVault)
					// log.Printf("Path from Vault: %v", pathFromVault)
					//Call a document Service instance
					aiService := utils.NewIDDocumentService(pathFromVault, secretFromVault, indexValue)
					aiService2 := utils.NewIDDocumentService(config.ServiceEndpoints, config.ServiceKeys, indexValue)
					log.Printf("Vault Values, %v", aiService)
					log.Printf("Config Vaules:%v", aiService2)
					persoFile, err := aiService.SelectDocument(indexValue)
					if err != nil {
						log.Fatalf("Error selecting document: %v", err)
					}
					//log.Printf(pathFromVault, secretFromVault, indexValue)
					aiServiceRequest, err := aiService.UploadDocument(persoFile)
					if err != nil {
						log.Fatalf("Error uploading document: %v", err)
					}
					log.Printf("Successfully uploaded document to : %v", aiServiceRequest)

					aiServiceResults, err := aiService.GetResults(aiServiceRequest)
					if err != nil {
						log.Printf("Error fetching results: %v", err)
					}
					log.Printf("Successfully fetched results: %v", aiServiceResults)
					aiServiceResultsParsed, err := aiService.ParseResults(aiServiceResults)
					if err != nil {
						log.Printf("Error parsing results: %v", err)
					}

					aiServiceResultsFormatted := aiService.FormatResults(aiServiceResultsParsed)

					// Pass the resulting map to the global variable to make it accessible
					// for the func that sends the citizen permit request to the server
					ProcessedResultMap = aiServiceResultsParsed
					Vendor = "Azure"

					// Layout
					resultsView.SetText("\nIntelligente Dokumentenverarbeitung powered by Azure\n\nPRÜFEN SIE AUF KORREKTHEIT UND BESTÄTIGEN SIE MIT 'ENTER'\n\n\n\n" + aiServiceResultsFormatted).SetTextColor(tcell.ColorYellow)
					pages.AddPage("AIResults", resultsView, true, true)
					pages.SwitchToPage("AIResults")
				})

			pages.AddPage("ProcessPage", modal, true, true)

		} else {
			modal := tview.NewModal().
				SetText("Dokument bereit zur Verarbeitung...!").
				AddButtons([]string{"Mit ENTER bestätigen"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					// Define the page that occurs when clicking the button
					resultsView := tview.NewTextView()
					resultsView.SetBorder(true).SetTitle("Ergebnisse der Dokumentenverarbeitung").SetTitleAlign(tview.AlignCenter)
					//resultsView.SetText("Bitte prüfen Sie die Ergebnisse der Dokumentenverarbeitung auf Korrektheit.")
					resultsView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
						if event.Key() == tcell.KeyEnter {
							pages.SwitchToPage("SmartPermitPage")
						}
						return event
					})

					// Process the document through AI service
					_, indexValue := smartDocumentForm.GetFormItemByLabel("Dokument Wählen").(*tview.DropDown).GetCurrentOption()
					aiService := utils.NewTextractClient(indexValue, config.AWSRegion)
					persoFile, err := aiService.SelectDocument(indexValue)
					if err != nil {
						log.Fatalf("Error selecting document: %v", err)
					}

					byteFile := utils.ConvertToTextractImage(persoFile)

					aiServiceRequest, err := aiService.AnalyzeID(byteFile)
					if err != nil {
						log.Fatalf("Error uploading document: %v", err)
					}

					resultsJSON, err := json.Marshal(aiServiceRequest)
					if err != nil {
						log.Fatalf("Error marshalling results: %v", err)
					}

					aiServiceResultsParsed, err := aiService.ParseResults(string(resultsJSON))
					if err != nil {
						log.Fatalf("Error parsing results: %v", err)
					}

					aiServiceResultsFormatted := aiService.FormatResults(aiServiceResultsParsed)

					// Pass the resulting map to the global variable to make it accessible
					// for the func that sends the citizen permit request to the server
					ProcessedResultMap = aiServiceResultsParsed
					Vendor = "AWS"

					// Layout
					resultsView.SetText("\nIntelligente Dokumentenverarbeitung powered by AWS\n\nPRÜFEN SIE AUF KORREKTHEIT UND BESTÄTIGEN SIE MIT 'ENTER'\n\n\n\n" + aiServiceResultsFormatted).SetTextColor(tcell.ColorYellow)
					pages.AddPage("AIResults", resultsView, true, true)
					pages.SwitchToPage("AIResults")
				})

			pages.AddPage("ProcessPage", modal, true, true)

		}
	})

	smartDocumentForm.AddButton("Zurück", func() { pages.SwitchToPage("CitizenLandingPage") })

	pages.AddPage("SmartDocumentPage", smartDocumentForm, true, false)
}

// Page Setup to process the AI-assisted documenta and then send the request
func SetupSmartPermitPage(app *tview.Application, pages *tview.Pages, config config.Config) {
	smartPermitForm := tview.NewForm()
	smartPermitForm.SetBorder(true).SetTitle("Erfassung von Pflichtangaben").SetTitleAlign(tview.AlignCenter)

	smartPermitForm.AddTextView("Intelligente Dokumentenverarbeitung powered by Azure/AWS", "Bitte befüllen Sie das untenstehende Formular. Die restlichen Angaben wurden durch den Dokumentenupload automatisch erfasst.	", 50, 0, false, false).SetBorder(true).SetBorderColor(tview.Styles.BorderColor)
	smartPermitForm.AddInputField("Datum der Ausstellung [dd:mm:yy]", "", 30, nil, nil)
	smartPermitForm.AddInputField("Geschlecht", "", 30, nil, nil)
	smartPermitForm.AddInputField("Nationalität", "", 30, nil, nil)
	smartPermitForm.AddInputField("Ausstellende Behörde [Amt]", "", 30, nil, nil)
	smartPermitForm.AddInputField("Antragstyp[Hobby/Gewerblich]", "", 30, nil, nil)
	smartPermitForm.AddInputField("Vorgesehene Nutzungsregion [Berlin/Brandenburg]", "", 30, nil, nil)
	smartPermitForm.AddInputField("E-Mail Adresse [max.mustermann@xxx.com]", "", 30, nil, nil)
	smartPermitForm.AddInputField("Telefonnummer[+49 12345678]", "", 30, nil, nil)
	smartPermitForm.AddButton("Antrag stellen", func() {
		if Vendor == "Azure" {
			citizenPermit := data.CreateMergedCitizenPermitFromService(smartPermitForm, ProcessedResultMap)

			// Call API handler to post citizen permit request
			requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[0])
			requestClient.TransferCitizenPermit(citizenPermit, APIAccessToken)
		} else {
			citizenPermit := data.CreateMergedCitizenPermitFromService2(smartPermitForm, ProcessedResultMap)

			// Call API handler to post citizen permit request
			requestClient := data.NewJSONTransferClient(config.ServerAddress, config.ServerPort, config.ServerAPIs[0])
			requestClient.TransferCitizenPermit(citizenPermit, APIAccessToken)

		}
	})

	// Buttons
	smartPermitForm.AddButton("Zurück", func() { pages.SwitchToPage("CitizenLandingPage") })

	// Finish Page Setup
	pages.AddPage("SmartPermitPage", smartPermitForm, true, false)
}
