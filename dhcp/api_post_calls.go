package dhcp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func ChangeDHCPSettings(active ServerSettings, settingsJson ServerConfig) error {
	jsonValue, err := json.Marshal(settingsJson)
	if err != nil {
		log.Println("JSON parse error: ", err)
		return err
	}

	err = apiCallPost(fmt.Sprintf("http://%s:%s/control/dhcp/set_config", active.Address, active.Port), active, jsonValue)
	if err != nil {
		log.Println("Error changing DHCP Settings")
		return err
	}
	return nil
}

func AddStaticLeases(target ServerSettings, lease StaticLeaseList) error {
	fmt.Printf("Adding/updating static leases to %s\n", target.Address)

	for _, x := range lease.StaticLeases {
		fmt.Printf("Adding Lease for Hostname: %s, IP: %s, MAC: %s \n", x.Hostname, x.IP, x.MAC)
		jsonValue, err := json.Marshal(x)
		if err != nil {
			log.Println("JSON parse error: ", err)
			return err
		}

		err = apiCallPost(fmt.Sprintf("http://%s:%s/control/dhcp/add_static_lease", target.Address, target.Port), target, jsonValue)
		if err != nil {
			log.Println("Error adding Static Lease")
			return err
		}
	}
	return nil
}

func apiCallPost(apiUrl string, active ServerSettings, jsonStruct []byte) error {
	url := fmt.Sprint(apiUrl)

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStruct))
	if err != nil {
		return err
	}
	req.Close = true

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}

	req.SetBasicAuth(active.Username, active.Password)

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return nil
}
