package main

import (
	"encoding/json"
	"log"
	"os"

	config "client/data"
	presentation "client/ui"

	"github.com/rivo/tview"
)

func main() {
	// Client Config
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config config.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	//Cli Instanciation
	app := tview.NewApplication()
	pages := tview.NewPages()

	// Setup Start Page
	presentation.SetupStartPage(app, pages)

	// Setup Citizen Pages
	presentation.SetupCitizenLandingPage(app, pages, config)

	presentation.SetupManualPermitPage(app, pages, config)

	presentation.SetupSmartPermitPage(app, pages, config)

	presentation.SetupSmartDocumentPage(app, pages, config)

	// Setup Admin Pages
	presentation.SetupAdminPage(app, pages, config)

	// Set the initial page to be the Start Page
	if err := app.SetRoot(pages, true).Run(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}

	//ID Document Test
	// aiService := utils.NewIDDocumentService(config.ServiceEndpoints, config.ServiceKeys, config.FilePaths[0])

	// file, err := aiService.SelectDocument(config.FilePaths[0])
	// if err != nil {
	// 	log.Fatalf("Error selecting document: %v", err)
	// }
	// log.Printf("Successfully selected document")

	// result, err := aiService.UploadDocument(file)
	// if err != nil {
	// 	log.Fatalf("Error uploading document: %v", err)
	// }
	// log.Printf("Successfully uploaded document to : %v", result)

	// aiResult, err := aiService.GetResults(result)
	// if err != nil {
	// 	log.Fatalf("Error getting results: %v", err)
	// }
	// aiService.ParseResults(aiResult)
}
