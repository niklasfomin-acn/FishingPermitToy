package auth

import (
	config "client/data"
	"log"

	"github.com/rivo/tview"
)

// Get the needed secret for whatever server from whatever secret manager that is configured
func GetSecretFromVault(hashicorp bool, azure bool, aws bool, config *config.Config) (string, string, error) {
	var endpoint, secret string
	if hashicorp {
		newClient, err := NewHashiCorpVaultClient(config)
		if err != nil {
			log.Fatalf("Error creating HashiCorp Vault client: %v", err)
		}
		err = newClient.ConnectToService()
		if err != nil {
			log.Fatalf("Error connecting to HashiCorp Vault service: %v", err)
		}
		log.Printf("Connected to HashiCorp Vault service")

		endpoint, err = newClient.GetSecret(config.SecretNames[0], config.SecretPaths[0])
		if err != nil {
			log.Fatalf("Error getting secret: %v", err)
		}
		log.Printf("Endpoint: %v", endpoint)

		secret, err = newClient.GetSecret(config.SecretNames[1], config.SecretPaths[0])
		if err != nil {
			log.Fatalf("Error getting secret: %v", err)
		}
		log.Printf("Secret: %v", secret)

	} else if azure {
		log.Printf("Azure Key Vault not implemented")
	} else if aws {
		log.Printf("AWS Secret Manager not implemented")
	}
	return endpoint, secret, nil
}

// Read the credentials from the login form
func GetCredentialsFromForm(form *tview.Form) (string, string) {
	username := form.GetFormItemByLabel("Nutzername").(*tview.InputField).GetText()
	password := form.GetFormItemByLabel("Password").(*tview.InputField).GetText()

	return username, password
}

func GetIdentityFromVault(config config.Config) (string, string, string, string, error) {
	var endpoint, secret, clientID, realm string
	newClient, err := NewHashiCorpVaultClient(&config)

	if err != nil {
		log.Fatalf("Error creating HashiCorp Vault client: %v", err)
	}
	err = newClient.ConnectToService()
	if err != nil {
		log.Fatalf("Error connecting to HashiCorp Vault service: %v", err)
	}
	log.Printf("Connected to HashiCorp Vault service")

	// Get the Keycloak Client ID
	clientID, err = newClient.GetSecret(config.SecretNames[3], config.SecretPaths[1])
	if err != nil {
		log.Fatalf("Error getting secret: %v", err)
	}
	log.Printf("Client ID: %v", clientID)

	// Get the Keycloak Endpoint
	endpoint, err = newClient.GetSecret(config.SecretNames[2], config.SecretPaths[1])
	if err != nil {
		log.Fatalf("Error getting secret: %v", err)
	}
	log.Printf("Endpoint: %v", endpoint)

	// Get the Keycloak Client Secret
	secret, err = newClient.GetSecret(config.SecretNames[4], config.SecretPaths[1])
	if err != nil {
		log.Fatalf("Error getting secret: %v", err)
	}
	log.Printf("Secret: %v", secret)

	// Get the Keycloak Realm
	realm, err = newClient.GetSecret(config.SecretNames[5], config.SecretPaths[1])
	if err != nil {
		log.Fatalf("Error getting secret: %v", err)
	}
	log.Printf("Realm: %v", realm)

	return clientID, endpoint, secret, realm, nil
}
