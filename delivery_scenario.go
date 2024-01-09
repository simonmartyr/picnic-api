package picnic

import (
	"strings"
)

// DeliveryScenario The Scenario content is a planned route for a delivery
// make use of DeliveryPosition to determine the current delivery location.
type DeliveryScenario struct {
	Version     int         `json:"version"`
	Destination Destination `json:"destination"`
	Driver      Driver      `json:"driver"`
	Scenario    []Scenario  `json:"scenario"`
	Vehicle     Vehicle     `json:"vehicle"`
}

type Destination struct {
	HouseNumber    int    `json:"house_number"`
	HouseNumberExt string `json:"house_number_ext"`
	PostCode       string `json:"post_code"`
	Street         string `json:"street"`
	City           string `json:"city"`
}

type Driver struct {
	Name     string `json:"name"`
	PhotoUrl string `json:"photo_url"`
	Quote    string `json:"quote"`
}

// Scenario
// Estimation of geolocation for a given time.
// Use the DeliveryPosition to find the current TimeStamp
type Scenario struct {
	TimeStamp int     `json:"ts"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
}

type Vehicle struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

// GetDeliveryScenario Query the current delivery information regarding an active delivery
//
// Method requires client to be authenticated
func (c *Client) GetDeliveryScenario(deliveryId string) (*DeliveryScenario, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if strings.TrimSpace(deliveryId) == "" {
		return nil, createError("GetDelivery requires a valid deliveryId string")
	}
	searchUrl := c.baseURL + "/deliveries/" + deliveryId + "/scenario"
	var deliveryScenario DeliveryScenario
	err := c.get(searchUrl, &deliveryScenario)
	if err != nil {
		return nil, err
	}
	return &deliveryScenario, nil
}
