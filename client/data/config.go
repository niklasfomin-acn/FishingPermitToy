package data

type Config struct {
	ServerAddress string   `json:"serverAddress"`
	ServerPort    string   `json:"serverPort"`
	SecretStore   string   `json:"secretStore"`
	ServerAPIs    []string `json:"serverAPIs"`
}
