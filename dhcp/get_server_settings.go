package dhcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetServerSettings(active ServerSettings, settings ServerConfig, leases StaticLeaseList) (ServerConfig, StaticLeaseList, error) {
	url := fmt.Sprintf("http://%s:%s/control/dhcp/status", active.Address, active.Port)

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print("Request failed!!!!!")
		log.Print(err)
		return settings, leases, err
	}
	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}

	if active.Username != "" && active.Password != "" {
		req.SetBasicAuth(active.Username, active.Password)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return settings, leases, err
	}

	if response.StatusCode != 200 {
		err = errors.New("HTTP status code " + strconv.Itoa(response.StatusCode))
		log.Printf("Error: %d", response.StatusCode)
		return settings, leases, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
		return settings, leases, err
	}

	defer response.Body.Close()

	err = json.Unmarshal(body, &settings)
	if err != nil {
		log.Println("JSON parse error: ", err)
		return settings, leases, err
	}
	err = json.Unmarshal(body, &leases)
	if err != nil {
		log.Println("JSON parse error: ", err)
		return settings, leases, err
	}

	return settings, leases, nil
}
