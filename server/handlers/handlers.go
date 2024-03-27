package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	types "rest-backend/types"
	"sync"
)

var (
	mu               sync.Mutex
	receivedPermits  = make([]types.Permit, 0)
	receivedCitizens = make([]types.Citizen, 0)
)

// Handle a newly received permit request
func HandlePermitRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var permit types.Permit
	err := json.NewDecoder(r.Body).Decode(&permit)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	receivedPermits = append(receivedPermits, permit)
	mu.Unlock()
	log.Printf("Received permit request: %+v\n", permit)

	w.WriteHeader(http.StatusCreated)

}

// Handle a newly received citizen request
func HandleCitizenRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var citizen types.Citizen
	err := json.NewDecoder(r.Body).Decode(&citizen)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	receivedCitizens = append(receivedCitizens, citizen)
	mu.Unlock()
	log.Printf("Received citizen request: %+v\n", citizen)

	w.WriteHeader(http.StatusCreated)

}
