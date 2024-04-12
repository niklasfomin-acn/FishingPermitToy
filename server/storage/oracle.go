package storage

import (
	"database/sql"
	"fmt"
	"log"
	"rest-backend/types"
)

type OracleStorage struct {
	db  *sql.DB
	uri string
}

func NewOracleStorage(uri string) (*OracleStorage, error) {
	db, err := sql.Open("goracle", uri)
	if err != nil {
		return nil, fmt.Errorf("error instanciating oracle db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging oracle db: %w", err)
	}

	os := &OracleStorage{
		db:  db,
		uri: uri,
	}

	err = os.CreateTable()
	if err != nil {
		return nil, fmt.Errorf("error creating table: %w", err)
	}

	return os, nil
}

func (os *OracleStorage) CreateTable() error {
	query := `
		CREATE TABLE citizenPermit (
			PassportNumber VARCHAR2(100) NOT NULL PRIMARY KEY,
            Surname VARCHAR2(100),
            GivenNames VARCHAR2(100),
            DateOfBirth VARCHAR2(100),
            PlaceOfBirth VARCHAR2(100),
            Gender VARCHAR2(10),
            Nationality VARCHAR2(100),
            DateOfIssue VARCHAR2(100),
            ExpiryDate VARCHAR2(100),
            IssuingAuthority VARCHAR2(100),
            PermitDate DATE,
            PermitLocation VARCHAR2(100),
            PermitType VARCHAR2(100),
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
