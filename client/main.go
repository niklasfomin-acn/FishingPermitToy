package main

import (
	"encoding/json"
	"log"
	"os"

	config "client/data"
	presentation "client/ui"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/rivo/tview"
)

func main() {
	//Client Config
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config config.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)
	if err != nil {
		log.Fatalf("Error creating session: %v", err)
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

	presentation.SetupSmartDocumentPage(app, pages, config, sess)

	// Setup Admin Pages
	presentation.SetupAdminPage(app, pages, config)

	// Set the initial page to be the Start Page
	if err := app.SetRoot(pages, true).Run(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}
}
