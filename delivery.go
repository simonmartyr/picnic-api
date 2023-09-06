package picnic

import (
	"strings"
)

type Delivery struct {
	Type               string            `json:"type"`
	Id                 string            `json:"id"`
	DeliveryId         string            `json:"delivery_id"`
	CreationTime       string            `json:"creation_time"`
	Slot               DeliverySlot      `json:"slot"`
	Eta2               DeliveryTime      `json:"eta_2"`
	Status             DeliveryStatus    `json:"status"`
	DeliveryTime       DeliveryTime      `json:"delivery_time"`
	Orders             []Order           `json:"orders"`
	ReturnedContainers []ReturnContainer `json:"returned_containers"`
	Parcels            []string          `json:"parcels"`
}

// GetDeliveries Query for all current or past deliveries. Optionally provide a filter of the list of DeliveryStatus
// to filter the deliveries by. The data returned is a summary, to get the complete data of a delivery use GetDelivery
//
// Method requires client to be authenticated
func (c *Client) GetDeliveries(filter []DeliveryStatus) (*[]Delivery, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if filter == nil {
		filter = []DeliveryStatus{}
	}
	searchUrl := c.baseURL + "/deliveries/summary"
	var deliveries []Delivery
	err := c.post(searchUrl, filter, &deliveries)
	if err != nil {
		return nil, err
	}
	return &deliveries, nil
}

// GetDelivery Query for the complete details of a particular delivery.
//
// Method requires client to be authenticated
func (c *Client) GetDelivery(deliveryId string) (*Delivery, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if strings.TrimSpace(deliveryId) == "" {
		return nil, createError("GetDelivery requires a valid deliveryId string")
	}
	searchUrl := c.baseURL + "/deliveries/" + deliveryId
	var delivery Delivery
	err := c.get(searchUrl, &delivery)
	if err != nil {
		return nil, err
	}
	return &delivery, nil
}
