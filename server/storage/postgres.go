/*
	GPT Experiment zum Untersuchungsaspekt Module / Datenbanken

Code Snippet zur Anbindung von AWS RDS Postgres an das Backend
Version: 3
Bemerkungen: Die Datenbankanbindung funktioniert. Allerdings behandelt der generierte Code nicht das
Szenario, in dem die Datenbanktabelle bereits existiert. Daher wirf der Server folgenden Fehler und stoppt.
2024/04/17 09:37:01 Error creating table: pq: relation "citizenpermit" already exists
2024/04/17 09:37:01 Table created successfully
2024/04/17 09:37:01 error creating table: pq: relation "citizenpermit" already exists
exit status 1

Der Fehler konnte durch die vorgschlagene Änderung im Code behoben werden. Der Server
startet und verarbeitet Anfragen korrekt.
*/
package storage

import (
	"database/sql"
	"fmt"
	"log"
	"rest-backend/types"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db  *sql.DB
	uri string
}

func NewPostgresStorage(uri string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("error instanciating postgres db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging postgres db: %w", err)
	}

	ps := &PostgresStorage{
		db:  db,
		uri: uri,
	}

	err = ps.CreateTable()
	if err != nil {
		return nil, fmt.Errorf("error creating table: %w", err)
	}

	return ps, nil
}

func (ps *PostgresStorage) CreateTable() error {
	query := `  
		CREATE TABLE IF NOT EXISTS citizenPermit (  
			PassportNumber VARCHAR(100) NOT NULL PRIMARY KEY,  
            Surname VARCHAR(100),  
            GivenNames VARCHAR(100),  
            DateOfBirth VARCHAR(100),  
            PlaceOfBirth VARCHAR(100),  
            Gender VARCHAR(10),  
            Nationality VARCHAR(100),  
            DateOfIssue VARCHAR(100),  
            ExpiryDate VARCHAR(100),  
            IssuingAuthority VARCHAR(100),  
            PermitDate DATE,  
            PermitLocation VARCHAR(100),  
            PermitType VARCHAR(100),  
            PermitState VARCHAR(100) NOT NULL  
        )  
    `
	_, err := ps.db.Exec(query)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return err
	}
	log.Println("Table created successfully")
	return nil
}

func (ps *PostgresStorage) SaveCitizenPermitRequest(cpr types.CitizenPermit) (interface{}, error) {
	query := `  
		INSERT INTO citizenPermit (  
			PassportNumber,  
            Surname,  
            GivenNames,  
            DateOfBirth,  
            PlaceOfBirth,  
            Gender,  
            Nationality,  
            DateOfIssue,  
            ExpiryDate,  
            IssuingAuthority,  
            PermitDate,  
            PermitLocation,  
            PermitType,  
            PermitState  
        ) VALUES (  
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14  
        )  
    `
	_, err := ps.db.Exec(query, cpr.PassportNumber, cpr.Surname, cpr.GivenNames, cpr.DateOfBirth, cpr.PlaceOfBirth, cpr.Gender, cpr.Nationality, cpr.DateOfIssue, cpr.ExpiryDate, cpr.IssuingAuthority, cpr.PermitDate, cpr.PermitLocation, cpr.PermitType, cpr.PermitState)
	if err != nil {
		log.Printf("Error saving citizen permit request: %v", err)
		return nil, err
	}
	log.Println("Citizen Permit Request saved successfully")
	return cpr, nil
}

func (ps *PostgresStorage) FetchCitizenPermitRequests() ([]types.CitizenPermit, error) {
	return nil, nil
}

func (ps *PostgresStorage) FetchProcessedCitizenPermitRequests() ([]types.CitizenPermit, error) {
	return nil, nil
}

func (ps *PostgresStorage) FetchPendingCitizenPermitRequests() ([]types.CitizenPermit, error) {
	return nil, nil
}

func (ps *PostgresStorage) FetchApprovedCitizenPermitRequests() ([]types.CitizenPermit, error) {
	return nil, nil
}

func (ps *PostgresStorage) FetchRejectedCitizenPermitRequests() ([]types.CitizenPermit, error) {
	return nil, nil
}

func (ps *PostgresStorage) FetchCitizenPermitRequestByID(id string) (types.CitizenPermit, error) {
	return types.CitizenPermit{}, nil
}

func (ps *PostgresStorage) ApproveCitizenPermitRequest(id string) error {
	return nil
}

func (ps *PostgresStorage) RejectCitizenPermitRequest(id string) error {
	return nil
}

func (ps *PostgresStorage) ViewPermitStatus(id string) (types.CitizenPermit, error) {
	return types.CitizenPermit{}, nil
}
