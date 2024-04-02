package types

import (
	"time"
)

type Citizen struct {
	PassportNumber   string `json:"passport_number"`
	Surname          string `json:"surname"`
	GivenNames       string `json:"given_names"`
	DateOfBirth      string `json:"date_of_birth"`
	PlaceOfBirth     string `json:"place_of_birth"`
	Gender           string `json:"gender"`
	Nationality      string `json:"nationality"`
	DateOfIssue      string `json:"date_of_issue"`
	ExpiryDate       string `json:"expiry_date"`
	IssuingAuthority string `json:"issuing_authority"`
}

type Permit struct {
	PermitDate     time.Time `json:"permit_date"`
	PermitLocation string    `json:"permit_location"`
	PermitType     string    `json:"permit_type"`
	PermitState    string    `json:"permit_state"`
}
