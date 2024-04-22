package utils

import (
	"image"
)

type DocumentService interface {

	// Get the document from the user
	SelectDocument(FilePath string) (file image.Image, err error)

	// Upload the document to the AI Service
	UploadDocument(File image.Image) (operationLocation string, err error)

	// Get the results from the AI Service
	GetResults(Endpoint string) (results string, err error)

	// Parse in the results from the AI service
	ParseResults(results string) (res map[string]interface{}, err error)

	// Format the parsed results
	FormatResults(results map[string]interface{}) string
}
