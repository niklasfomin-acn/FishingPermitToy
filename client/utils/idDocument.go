package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
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

func (doc *idDocument) UploadDocument(File image.Image) (JobId string, err error) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")}, // die Region anpassen
	)

	// Convert image.Image to []byte
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, File, nil)
	if err != nil {
		return JobId, err
	}

	// Create a Textract service client
	svc := textract.New(sess)

	// Call Textract
	result, err := svc.AnalyzeDocument(&textract.AnalyzeDocumentInput{
		Document: &textract.Document{
			Bytes: buf.Bytes(),
		},
		FeatureTypes: []*string{
			aws.String("FORMS"),
		},
	})

	if err != nil {
		return JobId, err
	}

	JobId = *result.JobId
	return JobId, nil
}

func (doc *idDocument) GetResults(jobId string) (results *textract.GetDocumentAnalysisOutput, err error) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")}, // die Region anpassen
	)
	svc := textract.New(sess)

	err = svc.GetDocumentAnalysisPages(&textract.GetDocumentAnalysisInput{
		JobId: aws.String(jobId),
	})

	return results, err
}

func (doc *idDocument) ParseResults(results *textract.GetDocumentAnalysisOutput) (res map[string]interface{}, err error) {
	for _, item := range results.Blocks {
		if *item.BlockType == "KEY_VALUE_SET" {
			if *item.EntityTypes[0] == "KEY" {
				key := *item.Text
				for _, valueID := range item.Value {
					for _, valueItem := range results.Blocks {
						if *valueItem.Id == *valueID {
							res[key] = *valueItem.Text
							break
						}
					}
				}
			}
		}
	}

	return res, nil
}

func (doc *idDocument) FormatResults(results map[string]interface{}) string {
	// 	var formattedResults string

	// 	for key, value := range results {
	// 		if key == "Address" {
	// 			if addressMap, ok := value.(map[string]interface{}); ok {
	// 				formattedResults += "				" + key + ": "
	// 				for addressKey, addressValue := range addressMap {
	// 					formattedResults += fmt.Sprintf("%s: %v, ", addressKey, addressValue)
	// 				}
	// 				formattedResults = strings.TrimSuffix(formattedResults, ", ") + "\n"
	// 			}
	// 		} else {
	// 			formattedResults += "								" + key + ": " + fmt.Sprint(value) + "\n"
	// 		}
	// 	}

	// 	return formattedResults
	// }

	var formattedResults string

	for key, value := range results {
		if key != "Address" {
			formattedResults += "								" + key + ": " + fmt.Sprint(value) + "\n"
		}
	}

	return formattedResults
}
