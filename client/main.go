package main

import (
	presentation "client/ui"
	"encoding/json"
	"log"
	"os"

	config "client/data"

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

	// Cli Instanciation
	app := tview.NewApplication()
	pages := tview.NewPages()

	// Setup Start Page
	presentation.SetupStartPage(app, pages)

	// Setup Citizen Pages
	presentation.SetupCitizenLandingPage(app, pages, config)

	presentation.SetupManualPermitPage(app, pages, config)

	// Setup Admin Pages
	presentation.SetupAdminPage(app, pages, config)

	presentation.SetupPermitPage(app, pages, config)

	// Set the initial page to be the Start Page
	if err := app.SetRoot(pages, true).Run(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}
}
