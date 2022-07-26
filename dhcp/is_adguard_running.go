package dhcp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type adguardStatus struct {
	Running bool `json:"enabled"`
}

// I would much rather if the API would return just the DHCP server status instead of sending the entire lease list with it.

func IsAdguardRunning(active ServerSettings) bool {
	var status adguardStatus = adguardStatus{Running: false}
	url := fmt.Sprintf("http://%s:%s/control/dhcp/status", active.Address, active.Port)

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}
	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}

	if active.Username != "" && active.Password != "" {
		req.SetBasicAuth(active.Username, active.Password)
	}

	response, err := client.Do(req)
	if err != nil {
		return false
	}

	if response.StatusCode != 200 {
		log.Printf("Error: %d", response.StatusCode)
		return false
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false
	}

	defer response.Body.Close()

	err = json.Unmarshal(body, &status)
	if err != nil {
		return false
	}
	return status.Running
}
