package auth

import (
	config "client/data"
	"log"
)

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
