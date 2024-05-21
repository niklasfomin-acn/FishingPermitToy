package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type jsonTransferClient struct {
	destAddr string
	destPort string
	destAPI  string
}

func NewJSONTransferClient(destAddr, destPort, destAPI string) *jsonTransferClient {
	return &jsonTransferClient{
		destAddr: destAddr,
		destPort: destPort,
		destAPI:  destAPI,
	}
}

// Citizen API Access
func (client *jsonTransferClient) TransferCitizenPermit(cp *CitizenPermit, token string) error {
	jsonData, err := json.Marshal(cp)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://"+client.destAddr+":"+client.destPort+client.destAPI, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send citizen permit data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-ok status from server: %v", resp.StatusCode)
	}

	return nil
}

func (client *jsonTransferClient) ViewCitizenPermitStatusByID(id string) error {
	return nil
}

// Admin API Access
func (client *jsonTransferClient) ViewAllCitizenPermitRequests() ([]CitizenPermit, error) {
	resp, err := http.Get("http://" + client.destAddr + ":" + client.destPort + client.destAPI)
	if err != nil {
		return []CitizenPermit{}, fmt.Errorf("failed to fetch all citizen permit requests: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []CitizenPermit{}, fmt.Errorf("received non-ok status from server: %v", resp.StatusCode)
	}

	var cp []CitizenPermit
	err = json.NewDecoder(resp.Body).Decode(&cp)
	if err != nil {
		log.Printf("Error decoding citizen permit data: %v", err)
		return []CitizenPermit{}, fmt.Errorf("failed to decode citizen permit data: %v", err)
	}
	log.Printf("Received citizen permit requests: %+v\n", cp)
	return cp, nil
}

func (client *jsonTransferClient) ViewPendingCitizenPermitRequests() ([]CitizenPermit, error) {
	resp, err := http.Get("http://" + client.destAddr + ":" + client.destPort + client.destAPI)
	if err != nil {
		return []CitizenPermit{}, fmt.Errorf("failed to fetch pending citizen permit requests: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []CitizenPermit{}, fmt.Errorf("received non-ok status from server: %v", resp.StatusCode)
	}

	var cp []CitizenPermit
	err = json.NewDecoder(resp.Body).Decode(&cp)
	if err != nil {
		log.Printf("Error decoding pending citizen permit data: %v", err)
		return []CitizenPermit{}, fmt.Errorf("failed to decode pending citizen permit data: %v", err)
	}
	log.Printf("Received pending citizen permit requests: %+v\n", cp)
	return cp, nil
}

func (client *jsonTransferClient) ViewProcessedCitizenPermitRequests() ([]CitizenPermit, error) {
	resp, err := http.Get("http://" + client.destAddr + ":" + client.destPort + client.destAPI)
	if err != nil {
		return []CitizenPermit{}, fmt.Errorf("failed to fetch processed citizen permit requests: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []CitizenPermit{}, fmt.Errorf("received non-ok status from server: %v", resp.StatusCode)
	}

	var cp []CitizenPermit
	err = json.NewDecoder(resp.Body).Decode(&cp)
	if err != nil {
		log.Printf("Error decoding processed citizen permit data: %v", err)
		return []CitizenPermit{}, fmt.Errorf("failed to decode processed citizen permit data: %v", err)
	}
	log.Printf("Received processed citizen permit requests: %+v\n", cp)
	return cp, nil
}

func (client *jsonTransferClient) ViewApprovedCitizenPermitRequests() ([]CitizenPermit, error) {
	resp, err := http.Get("http://" + client.destAddr + ":" + client.destPort + client.destAPI)
	if err != nil {
		return []CitizenPermit{}, fmt.Errorf("failed to fetch processed citizen permit requests: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []CitizenPermit{}, fmt.Errorf("received non-ok status from server: %v", resp.StatusCode)
	}

	var cp []CitizenPermit
	err = json.NewDecoder(resp.Body).Decode(&cp)
	if err != nil {
		log.Printf("Error decoding approved citizen permit data: %v", err)
		return []CitizenPermit{}, fmt.Errorf("failed to decode approved citizen permit data: %v", err)
	}
	log.Printf("Received approved citizen permit requests: %+v\n", cp)
	return cp, nil
}

func (client *jsonTransferClient) ViewRejectedCitizenPermitRequests() ([]CitizenPermit, error) {
	resp, err := http.Get("http://" + client.destAddr + ":" + client.destPort + client.destAPI)
	if err != nil {
		return []CitizenPermit{}, fmt.Errorf("failed to fetch rejected citizen permit requests: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []CitizenPermit{}, fmt.Errorf("received non-ok status from server: %v", resp.StatusCode)
	}

	var cp []CitizenPermit
	err = json.NewDecoder(resp.Body).Decode(&cp)
	if err != nil {
		log.Printf("Error decoding rejected citizen permit data: %v", err)
		return []CitizenPermit{}, fmt.Errorf("failed to decode rejected citizen permit data: %v", err)
	}
	log.Printf("Received rejected citizen permit requests: %+v\n", cp)
	return cp, nil

}

func (client *jsonTransferClient) ApproveCitizenPermitRequest(id string) error {
	req, err := http.NewRequest("POST", "http://"+client.destAddr+":"+client.destPort+client.destAPI+"?id="+id, nil)
	if err != nil {
		return fmt.Errorf("failed to create http request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send approval request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-ok status from server: %v", resp.StatusCode)
	}
	return nil
}

func (client *jsonTransferClient) RejectCitizenPermitRequest(id string) error {
	req, err := http.NewRequest("POST", "http://"+client.destAddr+":"+client.destPort+client.destAPI+"?id="+id, nil)
	if err != nil {
		return fmt.Errorf("failed to create http request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send rejection request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-ok status from server: %v", resp.StatusCode)
	}

	return nil
}
