package data

type Config struct {
	ServerAddress    string   `json:"serverAddress"`
	ServerPort       string   `json:"serverPort"`
	SecretStore      string   `json:"secretStore"`
	ServerAPIs       []string `json:"serverAPIs"`
	ServiceEndpoints string   `json:"serviceEndpoints"`
	ServiceKeys      string   `json:"serviceKeys"`
	FilePaths        []string `json:"FilePaths"`
	AWSRegion        string   `json:"AWSRegion"`
}
