package data

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (client *jsonTransferClient) TransferPermit(p *Permit) error {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://"+client.destAddr+":"+client.destPort+client.destAPI, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send permit data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-ok status from server: %v", resp.StatusCode)
	}
	return nil
}
