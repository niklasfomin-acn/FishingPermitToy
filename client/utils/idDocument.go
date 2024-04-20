package utils

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type idDocument struct {
	Endpoint string
	Key      string
	FilePath string
	File     image.Image
	client   http.Client
}

func NewIDDocumentService(Endpoint string, Key string, FilePath string) DocumentService {
	return &idDocument{
		Endpoint: Endpoint,
		Key:      Key,
		FilePath: FilePath,
		File:     nil,
		client:   http.Client{},
	}
}

func (doc *idDocument) SelectDocument(FilePath string) (File image.Image, err error) {
	document, err := os.Open(doc.FilePath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return nil, err
	}
	defer document.Close()

	File, _, err = image.Decode(document)
	if err != nil {
		log.Printf("Error decoding image: %v", err)
		return nil, err
	}
	return File, nil
}

func (doc *idDocument) UploadDocument(File image.Image) (operationLocation string, err error) {
	// Create a buffer to store the image
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, File, nil)
	if err != nil {
		return operationLocation, err
	}

	// Write the POST request
	req, err := http.NewRequest("POST", doc.Endpoint, buf)
	if err != nil {
		return operationLocation, err
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Ocp-Apim-Subscription-Key", doc.Key)

	resp, err := doc.client.Do(req)
	if err != nil {
		return operationLocation, err
	}
	defer resp.Body.Close()

	log.Printf("Response: %v", resp)
	operationLocation = resp.Header.Get("Operation-Location")
	return operationLocation, nil
}

func (doc *idDocument) GetResults(Endpoint string) (results string, err error) {
	for {
		req, err := http.NewRequest("GET", Endpoint, nil)
		if err != nil {
			return results, err
		}

		req.Header.Set("Ocp-Apim-Subscription-Key", doc.Key)
		response, err := doc.client.Do(req)
		if err != nil {
			return results, err
		}
		defer response.Body.Close()

		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return results, err
		}

		var resp http.Response
		err = json.Unmarshal(bodyBytes, &resp)
		if err != nil {
			return results, err
		}

		if resp.Status != "running" {
			results = string(bodyBytes)
			break
		} else {
			log.Printf("Response not ready yet: %v", response)
			time.Sleep(5 * time.Second)
		}
	}
	return results, nil
}

func (doc *idDocument) ConnectWithService() (client http.Client, err error) {
	client = http.Client{}

	return client, nil
}

func (doc *idDocument) ParseResults(results json.Decoder) (err error) {
	return nil
}
