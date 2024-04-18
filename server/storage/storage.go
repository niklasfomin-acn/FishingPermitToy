package storage

import "rest-backend/types"

type Storage interface {

	/* Citizen API storage procedures */

	// Save a CitizenPermitRequest
	SaveCitizenPermitRequest(cpr types.CitizenPermit) (interface{}, error)

	/* Admin API storage procedures */

	// Fetch all CitizenPermitRequests
	FetchCitizenPermitRequests() ([]types.CitizenPermit, error)

	// Fetch all processed CitizenPermitRequests
	FetchProcessedCitizenPermitRequests() ([]types.CitizenPermit, error)

	// Fetch all pending CitizenPermitRequests
	FetchPendingCitizenPermitRequests() ([]types.CitizenPermit, error)

	// Fetch all approved CitizenPermitRequests
	FetchApprovedCitizenPermitRequests() ([]types.CitizenPermit, error)

	// Fetch all rejected CitizenPermitRequests
	FetchRejectedCitizenPermitRequests() ([]types.CitizenPermit, error)

	// Fetch a CitizenPermitRequest by ID
	FetchCitizenPermitRequestByID(id string) (types.CitizenPermit, error)

	// Approve a CitizenPermitRequest
	ApproveCitizenPermitRequest(id string) error

	// Reject a CitizenPermitRequest
	RejectCitizenPermitRequest(id string) error
}
