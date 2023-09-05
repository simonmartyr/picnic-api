package picnic

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

func (c *Client) GetDeliveries(filter *[]DeliveryStatus) (*[]Delivery, error) {
	searchUrl := c.baseURL + "/deliveries/summary"
	var deliveries []Delivery
	err := c.post(searchUrl, filter, &deliveries)
	if err != nil {
		return nil, err
	}
	return &deliveries, nil
}

func (c *Client) GetDelivery(deliveryId string) (*Delivery, error) {
	searchUrl := c.baseURL + "/deliveries/" + deliveryId
	var delivery Delivery
	err := c.get(searchUrl, &delivery)
	if err != nil {
		return nil, err
	}
	return &delivery, nil
}
