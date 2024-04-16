package storage

import "rest-backend/types"

type Storage interface {

	// Admin Only!!! Fetch the Citizen Permit Requests
	// Save a CitizenPermitRequest
	SaveCitizenPermitRequest(cpr types.CitizenPermit) (interface{}, error)

	// Fetch all CitizenPermitRequests

	// Fetch all processed CitizenPermitRequests

	// Fetch all pending CitizenPermitRequests

	// Fetch all approved CitizenPermitRequests

	// Fetch all rejected CitizenPermitRequests

	// Fetch a CitizenPermitRequest by ID

	// Approve a CitizenPermitRequest

	// Reject a CitizenPermitRequest

}
