package storage

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"rest-backend/types"
)

type OracleStorage struct {
	db  *sql.DB
	url *url.URL
}

func NewOracleStorage(url *url.URL) (*OracleStorage, error) {
	db, err := sql.Open("goracle", url.String())
	if err != nil {
		return nil, fmt.Errorf("error instanciating oracle db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging oracle db: %w", err)
	}

	return &OracleStorage{
		db:  db,
		url: url,
	}, nil
}

func (os *OracleStorage) CreateTable() error {
	query := `
		CREATE TABLE citizenPermit (
			PassportNumber VARCHAR2(100) NOT NULL PRIMARY KEY,
            Surname VARCHAR2(100) NOT NULL,
            GivenNames VARCHAR2(100) NOT NULL,
            DateOfBirth VARCHAR2(100) NOT NULL,
            PlaceOfBirth VARCHAR2(100) NOT NULL,
            Gender VARCHAR2(10) NOT NULL,
            Nationality VARCHAR2(100) NOT NULL,
            DateOfIssue VARCHAR2(100) NOT NULL,
            ExpiryDate VARCHAR2(100) NOT NULL,
            IssuingAuthority VARCHAR2(100) NOT NULL,
            PermitDate DATE NOT NULL,
            PermitLocation VARCHAR2(100) NOT NULL,
            PermitType VARCHAR2(100) NOT NULL,
            PermitState VARCHAR2(100) NOT NULL
        )
    `
	_, err := os.db.Exec(query)
	if err != nil {
		log.Printf("Error creating table: %v", err)
	}
	log.Println("Table created successfully")
	return err
}

func (os *OracleStorage) SaveCitizenPermitRequest(cpr types.CitizenPermit) (interface{}, error) {
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
            :PassportNumber,
            :Surname,
            :GivenNames,
            :DateOfBirth,
            :PlaceOfBirth,
            :Gender,
            :Nationality,
			:DateOfIssue,
            :ExpiryDate,
            :IssuingAuthority,
            :PermitDate,
            :PermitLocation,
            :PermitType,
            :PermitState
        )
    `
	_, err := os.db.Exec(query, sql.Named("DateOfBirth", cpr.DateOfBirth), sql.Named("PlaceOfBirth", cpr.PlaceOfBirth), sql.Named("Gender", cpr.Gender), sql.Named("Nationality", cpr.Nationality), sql.Named("DateOfIssue", cpr.DateOfIssue), sql.Named("ExpiryDate", cpr.ExpiryDate), sql.Named("IssuingAuthority", cpr.IssuingAuthority), sql.Named("PermitDate", cpr.PermitDate), sql.Named("PermitLocation", cpr.PermitLocation), sql.Named("PermitType", cpr.PermitType), sql.Named("PermitState", cpr.PermitState))
	if err != nil {
		log.Printf("Error saving citizen permit request: %v", err)
		return nil, err
	}
	log.Println("Citizen Permit Request saved successfully")
	return cpr, nil
}
