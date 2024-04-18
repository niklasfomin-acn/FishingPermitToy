package storage

import (
	"fmt"
	"log"
	"rest-backend/types"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db  *sqlx.DB
	uri string
}

func NewPostgresStorage(uri string) (*PostgresStorage, error) {
	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("error instanciating postgres db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging postgres db: %w", err)
	}

	ps := PostgresStorage{
		db:  db,
		uri: uri,
	}

	err = ps.CreateTable()
	if err != nil {
		return nil, fmt.Errorf("error creating table: %w", err)
	}

	return &ps, nil
}

func (ps *PostgresStorage) CreateTable() error {
	if _, err := ps.db.Exec("DROP TABLE IF EXISTS citizenPermit"); err != nil {
		log.Printf("Error dropping table: %v", err)
	}
	query := `
		CREATE TABLE citizenPermit (
			PassportNumber VARCHAR(100) NOT NULL PRIMARY KEY,
            Surname VARCHAR(100),
            GivenNames VARCHAR(100),
            DateOfBirth VARCHAR(100) ,
            PlaceOfBirth VARCHAR(100) ,
            Gender VARCHAR(10) ,
            Nationality VARCHAR(100) ,
            DateOfIssue VARCHAR(100) ,
            ExpiryDate VARCHAR(100) ,
            IssuingAuthority VARCHAR(100) ,
            PermitDate DATE ,
            PermitLocation VARCHAR(100) ,
            PermitType VARCHAR(100) ,
            PermitState VARCHAR(100) 
        )
    `
	_, err := ps.db.Exec(query)
	if err != nil {
		log.Printf("Error creating table: %v", err)
	}
	log.Println("Table created successfully")
	return err
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
	_, err := ps.db.NamedExec(query, cpr)
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
