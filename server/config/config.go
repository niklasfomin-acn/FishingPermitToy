package config

type Config struct {
	ServerAddress             string            `json:"serverAddress"`
	Databases                 []string          `json:"databases"`
	Handlers                  map[string]string `json:"handlers"`
	HVaultAddress             string            `json:"HVaultAddress"`
	HVaultServiceEndpointPath string            `json:"ServiceEndpointPath"`
	HVaultAPIKeyPath          string            `json:"APIKeyPath"`
	SecretPaths               []string          `json:"SecretPaths"`
	SecretNames               []string          `json:"SecretNames"`
}
