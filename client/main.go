package main

import (
	presentation "client/ui"
	"log"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	pages := tview.NewPages()

	// Setup Start Page
	presentation.SetupStartPage(app, pages)

	// Setup Citizen Page
	presentation.SetupCitizenLandingPage(app, pages)

	presentation.SetupManualPermitPage(app, pages)

	presentation.SetupAdminPage(app, pages)

	// Set the initial page to be the Start Page
	if err := app.SetRoot(pages, true).Run(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}
}
