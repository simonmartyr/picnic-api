package picnic

import (
	"strings"
)

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

// GetCart retrieve the current cart contents of the authenticated user
//
// Method requires client to be authenticated
func (c *Client) GetCart() (*Order, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	cartUrl := c.baseURL + "/cart"
	var cart Order
	err := c.get(cartUrl, &cart)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

// AddToCart mutation method to add a quantity of a particular article to cart.
// The resulting Order contains the new state of the cart.
//
// Method requires client to be authenticated
func (c *Client) AddToCart(itemId string, count int) (*Order, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if strings.TrimSpace(itemId) == "" {
		return nil, createError("AddToCart requires a valid itemId string")
	}
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

// RemoveFromCart mutation method to remove a quantity of a particular article from the cart.
// The resulting Order contains the new state of the cart.
//
// Method requires client to be authenticated
func (c *Client) RemoveFromCart(itemId string, count int) (*Order, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if strings.TrimSpace(itemId) == "" {
		return nil, createError("RemoveFromCart requires a valid itemId string")
	}
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

// ClearCart mutation method to remove all articles from the cart.
// The resulting Order contains the new state of the cart.
//
// Method requires client to be authenticated
func (c *Client) ClearCart() (*Order, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	clearUrl := c.baseURL + "/cart/clear"
	var order = Order{}
	err := c.post(clearUrl, nil, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// SetDeliverySlot mutation method to select a delivery slot.
// The resulting Order contains the new state of the cart.
//
// Method requires client to be authenticated
func (c *Client) SetDeliverySlot(slotId string) (*Order, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if strings.TrimSpace(slotId) == "" {
		return nil, createError("SetDeliverySlot requires a valid slotId string")
	}
	clearUrl := c.baseURL + "/cart/set_delivery_slot"
	slot := SetSlot{SlotId: slotId}
	var order = Order{}
	err := c.post(clearUrl, slot, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
