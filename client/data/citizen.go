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
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
}

func CreateCitizenPermitFromForm(form *tview.Form) *CitizenPermit {
	permitDate := time.Now().Local()

	return &CitizenPermit{
		PassportNumber:   form.GetFormItemByLabel("Personalausweis-Nr.").(*tview.InputField).GetText(),
		Surname:          form.GetFormItemByLabel("Nachname").(*tview.InputField).GetText(),
		GivenNames:       form.GetFormItemByLabel("Vorname").(*tview.InputField).GetText(),
		DateOfBirth:      form.GetFormItemByLabel("Geburtsdatum [dd:mm:yy]").(*tview.InputField).GetText(),
		PlaceOfBirth:     form.GetFormItemByLabel("Geburtsort").(*tview.InputField).GetText(),
		Gender:           form.GetFormItemByLabel("Geschlecht").(*tview.InputField).GetText(),
		Nationality:      form.GetFormItemByLabel("Nationalität").(*tview.InputField).GetText(),
		DateOfIssue:      form.GetFormItemByLabel("Datum der Ausstellung [dd:mm:yy]").(*tview.InputField).GetText(),
		ExpiryDate:       form.GetFormItemByLabel("Gültig bis: [dd:mm:yy]").(*tview.InputField).GetText(),
		IssuingAuthority: form.GetFormItemByLabel("Ausstellende Behörde [Amt]").(*tview.InputField).GetText(),
		PermitDate:       permitDate,
		PermitLocation:   form.GetFormItemByLabel("Vorgesehene Nutzungsregion [Berlin/Brandenburg]").(*tview.InputField).GetText(),
		PermitType:       form.GetFormItemByLabel("Antragstyp[Hobby/Gewerblich]").(*tview.InputField).GetText(),
		PermitState:      "pending",
		Email:            form.GetFormItemByLabel("E-Mail Adresse [max.mustermann@xxx.com]").(*tview.InputField).GetText(),
		Phone:            form.GetFormItemByLabel("Telefonnummer[+49 12345678]").(*tview.InputField).GetText(),
	}
}

func CreateMergedCitizenPermitFromService(form *tview.Form, idData map[string]interface{}) *CitizenPermit {
	permitDate := time.Now().Local()

	return &CitizenPermit{
		PassportNumber:   idData["DocumentNumber"].(string),
		Surname:          idData["LastName"].(string),
		GivenNames:       idData["FirstName"].(string),
		DateOfBirth:      idData["DateOfBirth"].(string),
		PlaceOfBirth:     idData["PlaceOfBirth"].(string),
		Gender:           form.GetFormItemByLabel("Geschlecht").(*tview.InputField).GetText(),
		Nationality:      form.GetFormItemByLabel("Nationalität").(*tview.InputField).GetText(),
		DateOfIssue:      form.GetFormItemByLabel("Datum der Ausstellung [dd:mm:yy]").(*tview.InputField).GetText(),
		ExpiryDate:       idData["DateOfExpiration"].(string),
		IssuingAuthority: form.GetFormItemByLabel("Ausstellende Behörde [Amt]").(*tview.InputField).GetText(),
		PermitLocation:   form.GetFormItemByLabel("Vorgesehene Nutzungsregion [Berlin/Brandenburg]").(*tview.InputField).GetText(),
		PermitType:       form.GetFormItemByLabel("Antragstyp[Hobby/Gewerblich]").(*tview.InputField).GetText(),
		Email:            form.GetFormItemByLabel("E-Mail Adresse [max.mustermann@xxx.com]").(*tview.InputField).GetText(),
		Phone:            form.GetFormItemByLabel("Telefonnummer[+49 12345678]").(*tview.InputField).GetText(),
		PermitDate:       permitDate,
	}
}
