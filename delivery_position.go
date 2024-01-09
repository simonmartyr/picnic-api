package picnic

import (
	"strings"
)

// DeliveryPosition the current status of the delivery
// contains information regarding when the estimated time of arrival.
// ScenarioTs is a reference to the current geolocation within the DeliveryScenario.
type DeliveryPosition struct {
	Version            int       `json:"version"`
	ScenarioTs         int       `json:"scenario_ts"`
	Eta                int       `json:"eta"`
	EtaWindow          EtaWindow `json:"eta_window"`
	QueryInterval      int       `json:"query_interval"`
	ScenarioInProgress bool      `json:"scenario_in_progress"`
}

type EtaWindow struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// GetDeliveryPosition Query the current delivery position and estimated time of arrival
//
// Method requires client to be authenticated
func (c *Client) GetDeliveryPosition(deliveryId string) (*DeliveryPosition, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if strings.TrimSpace(deliveryId) == "" {
		return nil, createError("GetDelivery requires a valid deliveryId string")
	}
	searchUrl := c.baseURL + "/deliveries/" + deliveryId + "/position"
	var deliveryPosition DeliveryPosition
	err := c.get(searchUrl, &deliveryPosition)
	if err != nil {
		return nil, err
	}
	return &deliveryPosition, nil
}
