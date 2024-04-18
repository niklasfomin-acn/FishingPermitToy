package handlers

import (
	"net/http"
)

func GetHandlerFuncs(h *Handlers) map[string]func(w http.ResponseWriter, r *http.Request) {
	return map[string]func(w http.ResponseWriter, r *http.Request){
		"/SaveCitizenPermit":                 h.HandleCitizenPermitRequest,
		"/GetAllCitizenPermitRequests":       h.GetCitizenPermitRequests,
		"/GetProcessedCitizenPermitRequests": h.GetProcessedCitizenPermitRequests,
		"/GetPendingCitizenPermitRequests":   h.GetPendingCitizenPermitRequests,
		"/GetApprovedCitizenPermitRequests":  h.GetApprovedCitizenPermitRequests,
		"/GetRejectedCitizenPermitRequests":  h.GetRejectedCitizenPermitRequests,
		"/GetCitizenPermitRequestByID":       h.GetCitizenPermitRequestByID,
		"/ApproveCitizenPermitRequest":       h.ApproveCitizenPermitRequest,
		"/RejectCitizenPermitRequest":        h.RejectCitizenPermitRequest,
	}
}
