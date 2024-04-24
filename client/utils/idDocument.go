package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

type idDocument struct {
	Endpoint  string
	Key       string
	FilePaths string
	File      image.Image
	sess      *session.Session
}

func NewIDDocumentService(sess *session.Session, FilePaths string) DocumentService {
	return &idDocument{
		FilePaths: FilePaths,
		File:      nil,
		sess:      sess,
	}
}

func (doc *idDocument) SelectDocument() (File image.Image, err error) {
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
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, File, nil)
	if err != nil {
		return "", err
	}

	svc := textract.New(doc.sess)

	result, err := svc.DetectDocumentText(&textract.DetectDocumentTextInput{
		Document: &textract.Document{
			Bytes: buf.Bytes(),
		},
	})
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

func (doc *idDocument) GetResults(Endpoint string) (results string, err error) {
	/*
	   You could process the JobStatus from the Textract API for Asynchronous operations here
	   or print the output data if this function is used.
	   For now, it only returns an error
	*/
	return "", fmt.Errorf("function GetResults not implemented")
}

func (doc *idDocument) ParseResults(file string) (res map[string]interface{}, err error) {
	/*
	   Since we now return a string from UploadDocument,
	   we need to modify this function to parse the string into usable data.
	   Currently, this function returns an error
	*/
	return nil, fmt.Errorf("function ParseResults not implemented")
}

func (doc *idDocument) FormatResults(data map[string]interface{}) string {
	/*
	   This function is used to format the map[string]interface{} to your needs.
	   Here you would implement how the parsed data should be displayed or returned.
	   Currently, it returns an empty string.
	*/
	return ""
}
