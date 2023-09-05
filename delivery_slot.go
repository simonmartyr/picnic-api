package picnic

type DeliverySlot struct {
	SlotId               string `json:"slot_id"`
	HubId                string `json:"hub_id"`
	FcId                 string `json:"fc_id"`
	WindowStart          string `json:"window_start"`
	CutOffTime           string `json:"cut_off_time"`
	IsAvailable          bool   `json:"is_available"`
	Icon                 Icon   `json:"icon"`
	Selected             bool   `json:"selected"`
	Reserved             bool   `json:"reserved"`
	MinimumOrderValue    int    `json:"minimum_order_value"`
	UnavailabilityReason string `json:"unavailability_reason"`
}

type DeliverySlots struct {
	DeliverySlots []DeliverySlot `json:"delivery_slots"`
	SelectedSlot  SelectedSlot   `json:"selected_slot"`
}

func (c *Client) GetDeliverySlots() (*DeliverySlots, error) {
	searchUrl := c.baseURL + "/cart/delivery_slots"
	var deliverySlots DeliverySlots
	err := c.get(searchUrl, &deliverySlots)
	if err != nil {
		return nil, err
	}
	return &deliverySlots, nil
}
