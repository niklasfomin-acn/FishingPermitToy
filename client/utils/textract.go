package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

type TextractClient struct {
	svc       *textract.Textract
	FilePaths string
	File      image.Image
	Region    string
}

func NewTextractClient(FilePaths string, region string) *TextractClient {
	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return &TextractClient{
		svc:       textract.New(session),
		FilePaths: FilePaths,
		File:      nil,
	}
}

func (c *TextractClient) SelectDocument(FilePaths string) (File image.Image, err error) {
	document, err := os.Open(c.FilePaths)
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

func (c *TextractClient) AnalyzeID(input *textract.AnalyzeIDInput) (*textract.AnalyzeIDOutput, error) {

	req, output := c.svc.AnalyzeIDRequest(input)
	err := req.Send()
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	outputJSON, err := json.MarshalIndent(output, "", "		")
	if err != nil {
		log.Printf("Error marshalling output: %v", err)
		return nil, err
	}

	log.Printf("Output: %s", outputJSON)
	return output, nil
}

func ConvertToTextractImage(file image.Image) *textract.AnalyzeIDInput {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, file, nil)
	if err != nil {
		log.Fatalf("Error encoding image: %v", err)
	}
	imgBytes := buf.Bytes()

	input := &textract.AnalyzeIDInput{
		DocumentPages: []*textract.Document{
			{
				Bytes: imgBytes,
			},
		},
	}
	return input
}

// Dummy func to meet interface
func UploadDocument(File image.Image) (operationLocation string, err error) {
	return "", nil
}

// Dummy func to meet interface
func GetResults(Endpoint string) (results string, err error) {
	return "", nil
}

func (c *TextractClient) ParseResults(results string) (res map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(results), &res)
	if err != nil {
		log.Printf("Error unmarshalling results: %v", err)
		return res, nil
	}

	extractedData := make(map[string]interface{})

	if identityDocuments, ok := res["IdentityDocuments"].([]interface{}); ok {
		for _, document := range identityDocuments {
			if documentMap, ok := document.(map[string]interface{}); ok {
				if identityDocumentFields, ok := documentMap["IdentityDocumentFields"].([]interface{}); ok {
					for _, field := range identityDocumentFields {
						if fieldMap, ok := field.(map[string]interface{}); ok {
							if fieldType, ok := fieldMap["Type"].(map[string]interface{}); ok {
								if text, ok := fieldType["Text"].(string); ok {
									if valueDetection, ok := fieldMap["ValueDetection"].(map[string]interface{}); ok {
										if textValue, ok := valueDetection["Text"].(string); ok {
											extractedData[text] = textValue
											if normalizedValue, ok := valueDetection["NormalizedValue"].(map[string]interface{}); ok {
												if value, ok := normalizedValue["Value"].(string); ok {
													extractedData[text] = value
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return extractedData, nil
}

func (c TextractClient) FormatResults(results map[string]interface{}) string {
	var formattedResults string

	for key, value := range results {
		if key != "Address" {
			formattedResults += "								" + key + ": " + fmt.Sprint(value) + "\n"
		}
	}

	return formattedResults
}
