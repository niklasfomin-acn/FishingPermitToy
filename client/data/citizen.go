package data

import (
	"time"

	"github.com/rivo/tview"
)

// TODO: Fix the wrong types here
type CitizenPermit struct {
	PassportNumber   string    `json:"passport_number"`
	Surname          string    `json:"surname"`
	GivenNames       string    `json:"given_names"`
	DateOfBirth      string    `json:"date_of_birth"`
	PlaceOfBirth     string    `json:"place_of_birth"`
	Gender           string    `json:"gender"`
	Nationality      string    `json:"nationality"`
	DateOfIssue      string    `json:"date_of_issue"`
	ExpiryDate       string    `json:"expiry_date"`
	IssuingAuthority string    `json:"issuing_authority"`
	PermitDate       time.Time `json:"permit_date"`
	PermitLocation   string    `json:"permit_location"`
	PermitType       string    `json:"permit_type"`
	PermitState      string    `json:"permit_state"`
}

func CreateCitizenPermitFromForm(form *tview.Form) *CitizenPermit {
	permitDate := time.Now().Local()

	return &CitizenPermit{
		PassportNumber:   form.GetFormItemByLabel("Passport Number").(*tview.InputField).GetText(),
		Surname:          form.GetFormItemByLabel("Surname").(*tview.InputField).GetText(),
		GivenNames:       form.GetFormItemByLabel("Given Names").(*tview.InputField).GetText(),
		DateOfBirth:      form.GetFormItemByLabel("Date of Birth").(*tview.InputField).GetText(),
		PlaceOfBirth:     form.GetFormItemByLabel("Place of Birth").(*tview.InputField).GetText(),
		Gender:           form.GetFormItemByLabel("Gender").(*tview.InputField).GetText(),
		Nationality:      form.GetFormItemByLabel("Nationality").(*tview.InputField).GetText(),
		DateOfIssue:      form.GetFormItemByLabel("Date of Issue").(*tview.InputField).GetText(),
		ExpiryDate:       form.GetFormItemByLabel("Expiry Date").(*tview.InputField).GetText(),
		IssuingAuthority: form.GetFormItemByLabel("Issuing Authority").(*tview.InputField).GetText(),
		PermitDate:       permitDate,
		PermitLocation:   form.GetFormItemByLabel("Permit Location").(*tview.InputField).GetText(),
		PermitType:       form.GetFormItemByLabel("Permit Type").(*tview.InputField).GetText(),
	}
}
