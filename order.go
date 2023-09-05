package picnic

type Order struct {
	Type               string         `json:"type"`
	Id                 string         `json:"id"`
	Items              []OrderLine    `json:"items"`
	DeliverySlots      []DeliverySlot `json:"delivery_slots"`
	SelectedSlot       SelectedSlot   `json:"selected_slot"`
	TotalCount         int            `json:"total_count"`
	TotalPrice         int            `json:"total_price"`
	CheckoutTotalPrice int            `json:"checkout_total_price"`
	Mts                int            `json:"mts"`
	TotalSavings       int            `json:"total_savings"`
	TotalDeposit       int            `json:"total_deposit"`
	Cancellable        bool           `json:"cancellable"`
	CreationTime       string         `json:"creation_time"`
	Status             string         `json:"status"`
}

func (c *Client) GetCart() (*Order, error) {
	cartUrl := c.baseURL + "/cart"
	var cart Order
	err := c.get(cartUrl, &cart)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (c *Client) AddToCart(itemId string, count int) (*Order, error) {
	addUrl := c.baseURL + "/cart/add_product"
	toAdd := AddProductInput{
		ProductId: itemId,
		Count:     count,
	}
	var order = Order{}
	err := c.post(addUrl, toAdd, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (c *Client) RemoveFromCart(itemId string, count int) (*Order, error) {
	removeUrl := c.baseURL + "/cart/remove_product"
	toRemove := AddProductInput{
		ProductId: itemId,
		Count:     count,
	}
	var order = Order{}
	err := c.post(removeUrl, toRemove, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (c *Client) ClearCart() (*Order, error) {
	clearUrl := c.baseURL + "/cart/clear"
	var order = Order{}
	err := c.post(clearUrl, nil, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (c *Client) SetDeliverySlot(slotId string) (*Order, error) {
	clearUrl := c.baseURL + "/cart/set_delivery_slot"
	slot := SetSlot{SlotId: slotId}
	var order = Order{}
	err := c.post(clearUrl, slot, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
