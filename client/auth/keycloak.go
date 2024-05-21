package auth

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type KeycloakClient struct {
	ClientID     string
	Endpoint     string
	ClientSecret string
	Token        string
	httpClient   *http.Client
}

func NewKeycloakClient(ClientID string, Endpoint string, ClientSecret string) *KeycloakClient {
	return &KeycloakClient{
		ClientID:     ClientID,
		Endpoint:     Endpoint,
		ClientSecret: ClientSecret,
		httpClient:   &http.Client{},
	}
}

func (k *KeycloakClient) GetToken(username string, password interface{}, realm string) (verified bool, error error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", k.ClientID)
	data.Set("client_secret", k.ClientSecret)
	data.Set("username", username)
	data.Set("password", password.(string))

	req, err := http.NewRequest(http.MethodPost, k.Endpoint+"/realms/"+realm+"/protocol/openid-connect/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		verified = false
		return verified, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	log.Printf("Request: %v", req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		verified = false
		return verified, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result["access_token"] != nil {
		k.Token = result["access_token"].(string)
		verified = true
	} else {
		log.Printf("Error getting token: %v", result)
		verified = false
		return false, errors.New("failed to get token from Keycloak")
	}
	return true, nil
}
