package config

type Config struct {
	ServerAddress string            `json:"serverAddress"`
	Databases     []string          `json:"databases"`
	Handlers      map[string]string `json:"handlers"`
}
