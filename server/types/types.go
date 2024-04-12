package types

import (
	"time"
)

type CitizenPermit struct {
	PassportNumber   string    `json:"passport_number" db:"PassportNumber"`
	Surname          string    `json:"surname" db:"Surname"`
	GivenNames       string    `json:"given_names" db:"GivenNames"`
	DateOfBirth      string    `json:"date_of_birth" db:"DateOfBirth"`
	PlaceOfBirth     string    `json:"place_of_birth" db:"PlaceOfBirth"`
	Gender           string    `json:"gender" db:"Gender"`
	Nationality      string    `json:"nationality" db:"Nationality"`
	DateOfIssue      string    `json:"date_of_issue" db:"DateOfIssue"`
	ExpiryDate       string    `json:"expiry_date" db:"ExpiryDate"`
	IssuingAuthority string    `json:"issuing_authority" db:"IssuingAuthority"`
	PermitDate       time.Time `json:"permit_date" db:"PermitDate"`
	PermitLocation   string    `json:"permit_location" db:"PermitLocation"`
	PermitType       string    `json:"permit_type" db:"PermitType"`
	PermitState      string    `json:"permit_state" db:"PermitState"`
}
