package auth

import (
	"fmt"
	"log"
	"rest-backend/config"

	"github.com/dgrijalva/jwt-go"
)

type RbacClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

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

		endpoint, err = newClient.GetSecret(config.SecretNames[0], config.SecretPaths[1])
		if err != nil {
			log.Fatalf("Error getting secret: %v", err)
		}
		log.Printf("Endpoint: %v", endpoint)

		secret, err = newClient.GetSecret(config.SecretNames[1], config.SecretPaths[1])
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

func ValidateCitizenToken(token string) (bool, error) {
	parser := &jwt.Parser{
		SkipClaimsValidation: true,
	}

	parsedToken, _, err := parser.ParseUnverified(token, &RbacClaims{})
	if err != nil {
		return false, err
	}

	if claims, ok := parsedToken.Claims.(*RbacClaims); ok {
		if claims.Role == "citizen" {
			return true, nil
		}
		return false, fmt.Errorf("role is not valid")
	}
	return false, fmt.Errorf("token is not valid")
}

func ValidateAdminToken(token string) (bool, error) {
	parser := &jwt.Parser{
		SkipClaimsValidation: true,
	}

	parsedToken, _, err := parser.ParseUnverified(token, &RbacClaims{})
	if err != nil {
		return false, err
	}

	if claims, ok := parsedToken.Claims.(*RbacClaims); ok {
		if claims.Role == "admin" {
			return true, nil
		}
		return false, fmt.Errorf("role is not valid")
	}
	return false, fmt.Errorf("token is not valid")
}
