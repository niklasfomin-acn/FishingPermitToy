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
	presentation.SetupStartPage(app, pages, config)

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

	// // Textract Test
	// textractService := utils.NewTextractClient(config.FilePaths[0], config.AWSRegion)

	// file, err := textractService.SelectDocument(config.FilePaths[0])
	// if err != nil {
	// 	log.Fatalf("Error selecting document: %v", err)
	// }
	// log.Printf("Successfully selected document")

	// // Convert the image to a byte array
	// buf := new(bytes.Buffer)
	// err = jpeg.Encode(buf, file, nil)
	// if err != nil {
	// 	log.Fatalf("Error encoding image: %v", err)
	// }
	// imgBytes := buf.Bytes()

	// input := &textract.AnalyzeIDInput{
	// 	DocumentPages: []*textract.Document{
	// 		{
	// 			Bytes: imgBytes,
	// 			//DocumentType: aws.String("ID_DOCUMENT"),
	// 		},
	// 	},
	// }
	// results, err := textractService.AnalyzeID(input)
	// if err != nil {
	// 	log.Fatalf("Error analyzing document: %v", err)
	// }
	// resultsJSON, err := json.Marshal(results)
	// if err != nil {
	// 	log.Fatalf("Error marshalling results: %v", err)
	// }
	// final, err := textractService.ParseResults(string(resultsJSON))
	// if err != nil {
	// 	log.Fatalf("Error parsing results: %v", err)
	// }
	// log.Printf("Results: %v", final)

	// finalString := textractService.FormatResults(final)
	// log.Printf("Formatted Results: %v", finalString)

	// // Secret Manager Test
	// vault, err := auth.NewHashiCorpVaultClient(&config)
	// if err != nil {
	// 	log.Fatalf("Error creating vault client: %v", err)
	// }
	// err = vault.ConnectToService()
	// if err != nil {
	// 	log.Fatalf("Error connecting to vault service: %v", err)
	// }
	// log.Printf("Connected to vault service")

	// secret, err := vault.GetSecret(config.SecretNames[0], config.SecretPaths[0])
	// if err != nil {
	// 	log.Fatalf("Error getting secret: %v", err)
	// }
	// log.Printf("Secret: %v", secret)

	// secret2, err := vault.GetSecret(config.SecretNames[1], config.SecretPaths[0])
	// if err != nil {
	// 	log.Fatalf("Error getting secret: %v", err)
	// }
	// log.Printf("Secret: %v", secret2)

	// This is the test for the helper function. Depending on the call on the definition in config.json the different
	// vaults can be selected.
	// testEndpoint, testKey, err := auth.GetSecretFromVault(true, false, false, &config)
	// if err != nil {
	// 	log.Fatalf("Error getting secret: %v", err)
	// }
	// log.Printf("Path: %v", testEndpoint)
	// log.Print("Key: ", testKey)

	// // Keycloak Test
	// keycloak := auth.NewKeycloakClient(config.KeycloakClientID, config.KeycloakEndpoint, config.KeycloakClientSecret)
	// log.Printf("Keycloak Client Initialized: %v", keycloak)

	// token, err := keycloak.GetToken("frank", "my-new-password", config.KeycloakRealm)
	// if err != nil {
	// 	log.Fatalf("Error getting token: %v", token)
	// }
	// log.Printf("Token: %v", keycloak.Token)
}
