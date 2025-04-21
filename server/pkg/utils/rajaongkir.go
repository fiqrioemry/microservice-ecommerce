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

type ShippingOption struct {
	Service     string  `json:"service"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
	ETD         string  `json:"etd"`
}

type RajaOngkirCostResponse struct {
	RajaOngkir struct {
		Results []struct {
			Costs []struct {
				Service     string `json:"service"`
				Description string `json:"description"`
				Cost        []struct {
					Value int    `json:"value"`
					ETD   string `json:"etd"`
				} `json:"cost"`
			} `json:"costs"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

func FetchShippingCost(originID, destinationID, weight int, courier string) ([]ShippingOption, error) {
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
		return nil, err
	}
	req.Header.Set("key", apiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var response RajaOngkirCostResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if len(response.RajaOngkir.Results) == 0 || len(response.RajaOngkir.Results[0].Costs) == 0 {
		return nil, errors.New("no shipping cost returned")
	}

	var options []ShippingOption
	for _, cost := range response.RajaOngkir.Results[0].Costs {
		if len(cost.Cost) > 0 {
			options = append(options, ShippingOption{
				Service:     cost.Service,
				Description: cost.Description,
				Cost:        float64(cost.Cost[0].Value),
				ETD:         cost.Cost[0].ETD,
			})
		}
	}

	return options, nil
}
