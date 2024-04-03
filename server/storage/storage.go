package storage

import "rest-backend/types"

type Storage interface {

	// Admin Only!!! Fetch the Citizen Permit Requests
	FetchAll() ([]types.CitizenPermit, error)

	// Save a CitizenPermitRequest into db of choice (choose in config.json)
	SaveCitizenPermitRequest(cpr types.CitizenPermit) (interface{}, error)

	// Change permit status

}
