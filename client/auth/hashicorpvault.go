package auth

import (
	config "client/data"
	"fmt"
	"log"

	vaultapi "github.com/hashicorp/vault/api"
)

type HashiCorpVaultClient struct {
	Endpoint    string `json:"HVaultAddress"`
	Client      *vaultapi.Client
	SecretPaths []string `json:"SecretPaths"`
	SecretNames []string `json:"SecretNames"`
}

func NewHashiCorpVaultClient(config *config.Config) (*HashiCorpVaultClient, error) {
	return &HashiCorpVaultClient{
		Endpoint:    config.HVaultAddress,
		SecretPaths: config.SecretPaths,
		SecretNames: config.SecretNames,
	}, nil
}

func (h *HashiCorpVaultClient) ConnectToService() error {
	//ctx := context.Background()

	config := &vaultapi.Config{
		Address: h.Endpoint,
	}

	client, err := vaultapi.NewClient(config)
	if err != nil {
		log.Printf("Error creating vault client: %v", err)
		return err
	}

	clientCheck, err := client.Sys().Health()
	if err != nil {
		log.Printf(h.Endpoint)
		log.Printf("Error checking vault health: %v", err)
		return err
	}
	if clientCheck != nil && clientCheck.Initialized {
		log.Printf("Vault client initialized")
	} else {
		log.Fatal("Vault client not initialized")
	}

	h.Client = client

	return nil
}

func (h *HashiCorpVaultClient) GetSecret(name string, path string) (string, error) {
	secretClient := h.Client

	secret, err := secretClient.Logical().Read(path)
	log.Printf("Path: %v", path)
	if err != nil {
		log.Printf("Error reading secret: %v", err)
		return "", err
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok || data == nil {
		return "", fmt.Errorf("'data' field not found in secret data or is not a map")
	}

	value, ok := secret.Data["data"].(map[string]interface{})[name]
	if !ok || value == nil {
		return "", fmt.Errorf(name, ": is not found in secret data")
	}

	valueStr, ok := value.(string)
	if !ok {
		return "", fmt.Errorf(name, " :is not a string")
	}

	return valueStr, nil
}
