package picnic

type DeliverySlot struct {
	SlotId                  string `json:"slot_id"`
	HubId                   string `json:"hub_id"`
	FcId                    string `json:"fc_id"`
	WindowStart             string `json:"window_start"`
	WindowEnd               string `json:"window_end"`
	CutOffTime              string `json:"cut_off_time"`
	IsAvailable             bool   `json:"is_available"`
	Icon                    Icon   `json:"icon"`
	Selected                bool   `json:"selected"`
	Reserved                bool   `json:"reserved"`
	MinimumOrderValue       int    `json:"minimum_order_value"`
	UnavailabilityReason    string `json:"unavailability_reason"`
	WindowPresentationColor string `json:"window_presentation_color"`
}

type DeliverySlots struct {
	DeliverySlots []DeliverySlot `json:"delivery_slots"`
	SelectedSlot  SelectedSlot   `json:"selected_slot"`
}

// GetDeliverySlots Retrieves all slots for the next 7 days for the authenticated user
// A more direct alternative to using GetCart and accessing the data from the Order.
//
// Method requires client to be authenticated
func (c *Client) GetDeliverySlots() (*DeliverySlots, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	searchUrl := c.baseURL + "/cart/delivery_slots"
	var deliverySlots DeliverySlots
	err := c.get(searchUrl, &deliverySlots)
	if err != nil {
		return nil, err
	}
	return &deliverySlots, nil
}
