package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type idDocument struct {
	Endpoint  string
	Key       string
	FilePaths string
	File      image.Image
	client    http.Client
}

func NewIDDocumentService(Endpoint string, Key string, FilePaths string) DocumentService {
	return &idDocument{
		Endpoint:  Endpoint,
		Key:       Key,
		FilePaths: FilePaths,
		File:      nil,
		client:    http.Client{},
	}
}

func (doc *idDocument) SelectDocument(FilePaths string) (File image.Image, err error) {
	document, err := os.Open(doc.FilePaths)
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
	log.Printf("Operation Location: %v", operationLocation)
	return operationLocation, nil
}

func (doc *idDocument) GetResults(Endpoint string) (results string, err error) {
	for {
		req, err := http.NewRequest("GET", doc.Endpoint, nil)
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

func (doc *idDocument) ParseResults(results string) (map[string]interface{}, error) {
	var res map[string]interface{}
	err := json.Unmarshal([]byte(results), &res)
	if err != nil {
		log.Printf("Error unmarshalling results: %v", err)
		return res, nil
	}

	extractedData := make(map[string]interface{})

	if analyzeResult, ok := res["analyzeResult"].(map[string]interface{}); ok {
		if documents, ok := analyzeResult["documents"].([]interface{}); ok {
			for _, document := range documents {
				if docMap, ok := document.(map[string]interface{}); ok {
					if fields, ok := docMap["fields"].(map[string]interface{}); ok {
						extractedData = make(map[string]interface{})
						for key, value := range fields {
							if key == "FirstName" || key == "LastName" || key == "PlaceOfBirth" {
								if valMap, ok := value.(map[string]interface{}); ok {
									extractedData[key] = valMap["content"]
								}
							} else if key == "DateOfBirth" || key == "DateOfExpiration" {
								if valMap, ok := value.(map[string]interface{}); ok {
									extractedData[key] = valMap["valueDate"]
								}
							} else if key == "Address" {
								if valMap, ok := value.(map[string]interface{}); ok {
									extractedData[key] = valMap["valueAddress"]
								}
							} else if key == "DocumentNumber" || key == "EyeColor" || key == "Height" {
								if valMap, ok := value.(map[string]interface{}); ok {
									extractedData[key] = valMap["valueString"]
								}

							}

						}

						//fmt.Println(extractedData)
					}
				}
			}
		}

	}
	//extractedMapToString := fmt.Sprintf("%v", extractedData)
	//log.Printf("Extracted Data: %v", extractedMapToString)
	//log.Printf("Extracted Data: %v", string(jsonStr))
	return extractedData, nil
}

func (doc *idDocument) FormatResults(results map[string]interface{}) string {
	var formattedResults string

	for key, value := range results {
		if key != "Address" {
			formattedResults += "" + key + ": " + fmt.Sprint(value) + "\n"
		}
	}

	return formattedResults
}
