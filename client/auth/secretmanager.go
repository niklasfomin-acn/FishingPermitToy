package auth

type SecretManager interface {

	// Get connection to the secrets service
	ConnectToService() error

	// Retrieve a secret for the document AI service
	GetSecret(name string, path string) (string, error)
}
