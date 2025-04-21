package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RajaOngkirCostResponse struct {
	RajaOngkir struct {
		Results []struct {
			Costs []struct {
				Cost []struct {
					Value int `json:"value"`
				} `json:"cost"`
			} `json:"costs"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

func FetchShippingCost(originID, destinationID, weight int, courier string) (float64, error) {
	apiKey := os.Getenv("RAJAONGKIR_API_KEY")

	payload := map[string]string{
		"origin":      fmt.Sprintf("%d", originID),
		"destination": fmt.Sprintf("%d", destinationID),
		"weight":      fmt.Sprintf("%d", weight),
		"courier":     courier,
	}
	data := ""
	for k, v := range payload {
		data += fmt.Sprintf("%s=%s&", k, v)
	}

	req, err := http.NewRequest("POST", "https://api.rajaongkir.com/starter/cost", bytes.NewBufferString(data))
	if err != nil {
		return 0, err
	}
	req.Header.Set("key", apiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var response RajaOngkirCostResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return 0, err
	}

	if len(response.RajaOngkir.Results) == 0 || len(response.RajaOngkir.Results[0].Costs) == 0 {
		return 0, errors.New("no shipping cost returned")
	}

	return float64(response.RajaOngkir.Results[0].Costs[0].Cost[0].Value), nil
}
