package data

type Config struct {
	ServerAddress             string   `json:"serverAddress"`
	ServerPort                string   `json:"serverPort"`
	SecretStore               string   `json:"secretStore"`
	ServerAPIs                []string `json:"serverAPIs"`
	ServiceEndpoints          string   `json:"serviceEndpoints"`
	ServiceKeys               string   `json:"serviceKeys"`
	FilePaths                 []string `json:"FilePaths"`
	AWSRegion                 string   `json:"AWSRegion"`
	HVaultAddress             string   `json:"HVaultAddress"`
	HVaultServiceEndpointPath string   `json:"ServiceEndpointPath"`
	HVaultAPIKeyPath          string   `json:"APIKeyPath"`
	SecretPaths               []string `json:"SecretPaths"`
	SecretNames               []string `json:"SecretNames"`
	VaultProvider             string   `json:"VaultProvider"`
	KeycloakClientID          string   `json:"KeycloakClientID"`
	KeycloakEndpoint          string   `json:"KeycloakEndpoint"`
	KeycloakClientSecret      string   `json:"KeycloakClientSecret"`
	KeycloakRealm             string   `json:"KeycloakRealm"`
}
