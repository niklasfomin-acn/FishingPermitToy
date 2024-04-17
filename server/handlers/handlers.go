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

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResults)
	w.WriteHeader(http.StatusOK)

}

// Get all processed citizen permit requests

// Get all pending citizen permit requests

// Get all approved citizen permit requests

// Get all rejected citizen permit requests

// Get a citizen permit request by ID

// Approve a citizen permit request

// Reject a citizen permit request

// Notify a citizen on permit status
