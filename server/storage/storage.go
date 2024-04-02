package storage

import "rest-backend/types"

type Storage interface {

	// Get all entries

	// Save permit
	SavePermitRequest(pr types.Permit) (interface{}, error)

	// Save citizen request
	SaveCitizenRequest(cd types.Citizen) (interface{}, error)

	// Change permit status

}
