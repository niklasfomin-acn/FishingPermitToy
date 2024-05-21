package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-backend/storage"
	types "rest-backend/types"
	"sync"
)

type Handlers struct {
	Store storage.Storage
}

func New(store storage.Storage) *Handlers {
	return &Handlers{Store: store}
}

var (
	mu                            sync.Mutex
	receivedCitizenPermitRequests = make([]types.CitizenPermit, 0)
)

// Store a newly received citizen permit request
func (h *Handlers) HandleCitizenPermitRequest(w http.ResponseWriter, r *http.Request) {
	// validCitizenToken, err := auth.ValidateCitizenToken(token)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// if validCitizenToken == false {
	// 	http.Error(w, "Invalid token", http.StatusUnauthorized)
	// 	return
	// }

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var citzenPermit types.CitizenPermit
	err := json.NewDecoder(r.Body).Decode(&citzenPermit)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return

	}

	mu.Lock()
	receivedCitizenPermitRequests = append(receivedCitizenPermitRequests, citzenPermit)
	mu.Unlock()
	log.Printf("Received citizen permit request: %+v\n", receivedCitizenPermitRequests)

	insertPermitID, err := h.Store.SaveCitizenPermitRequest(citzenPermit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Citizen Permit Request Archived",
		"permit_id": insertPermitID,
	})

}

// Get all citizen permit requests
func (h *Handlers) GetCitizenPermitRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	allCitizenPermitRequests, err := h.Store.FetchCitizenPermitRequests()
	if err != nil {
		log.Printf("Error fetching citizen permit requests from database: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResults, err := json.Marshal(allCitizenPermitRequests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Successfully Served All Citizen Permit Requests")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResults)
}

// Get all processed citizen permit requests

func (h *Handlers) GetProcessedCitizenPermitRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	processedCitizenPermitRequests, err := h.Store.FetchProcessedCitizenPermitRequests()
	if err != nil {
		log.Printf("Error fetching processed citizen permit requests from database: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResults, err := json.Marshal(processedCitizenPermitRequests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Successfully Served All Processed Citizen Permit Requests")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResults)
}

// Get all pending citizen permit requests
func (h *Handlers) GetPendingCitizenPermitRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	pendingCitizenPermitRequests, err := h.Store.FetchPendingCitizenPermitRequests()
	if err != nil {
		log.Printf("Error fetching processed citizen permit requests from database: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResults, err := json.Marshal(pendingCitizenPermitRequests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Successfully Served All Pending Citizen Permit Requests")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResults)
}

// Get all approved citizen permit requests
func (h *Handlers) GetApprovedCitizenPermitRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	approvedCitizenPermitRequests, err := h.Store.FetchApprovedCitizenPermitRequests()
	if err != nil {
		log.Printf("Error fetching processed citizen permit requests from database: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResults, err := json.Marshal(approvedCitizenPermitRequests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Successfully Served All Approved Citizen Permit Requests")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResults)
}

// Get all rejected citizen permit requests
func (h *Handlers) GetRejectedCitizenPermitRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	rejectedCitizenPermitRequests, err := h.Store.FetchRejectedCitizenPermitRequests()
	if err != nil {
		log.Printf("Error fetching processed citizen permit requests from database: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResults, err := json.Marshal(rejectedCitizenPermitRequests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Successfully Served All Rejected Citizen Permit Requests")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResults)
}

// Get a citizen permit request by ID
func (h *Handlers) GetCitizenPermitRequestByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	permitID, err := h.Store.FetchCitizenPermitRequestByID(id)
	if err != nil {
		log.Printf("Error fetching Permit Request by ID: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResults, err := json.Marshal(permitID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Successfully Served Citizen Permit Request by ID")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResults)

}

// Approve a citizen permit request
func (h *Handlers) ApproveCitizenPermitRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	err := h.Store.ApproveCitizenPermitRequest(id)
	if err != nil {
		log.Printf("Error approving citizen permit request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		log.Printf("Citizen Permit Request Approved")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Citizen Permit Request Approved",
		})
	}
}

// Reject a citizen permit request
func (h *Handlers) RejectCitizenPermitRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	err := h.Store.RejectCitizenPermitRequest(id)
	if err != nil {
		log.Printf("Error rejecting citizen permit request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Citizen Permit Request Rejected",
		})
	}
}

// Notify a citizen on permit status
