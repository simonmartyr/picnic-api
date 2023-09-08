package picnic

import (
	"errors"
	"strings"
)

// CheckoutStart the Mts value should match the same value as the order
// The OosArticleIds likely refers to 'out of stock article ids'
// ResolveKey in the event a checkout requires additional verification the key must be included
// For example, the purchase of alcohol requires the value 'age_verified' to be present
type CheckoutStart struct {
	Mts           string `json:"mts"`
	OosArticleIds any    `json:"oos_article_ids"`
	ResolveKey    string `json:"resolve_key,omitempty"`
}

type Checkout struct {
	OrderId           string             `json:"order_id"`
	Address           Address            `json:"address"`
	DeliverySlots     []DeliverySlot     `json:"delivery_slots"`
	TotalPrice        int                `json:"total_price"`
	TransactionExpiry string             `json:"transaction_expiry"`
	TotalCount        int                `json:"total_count"`
	TotalDeposit      int                `json:"total_deposit"`
	TotalSavings      int                `json:"total_savings"`
	DepositBreakdown  []DepositBreakdown `json:"deposit_breakdown"`
}

type DepositBreakdown struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
	Count int    `json:"count"`
}

func (c *Client) StartCheckout(mts string) (*Checkout, *CheckoutError) {
	if !c.IsAuthenticated() {
		return nil, wrapCheckoutError(authenticationError())
	}
	if strings.TrimSpace(mts) == "" {
		return nil, wrapCheckoutError(createError("StartCheckout requires a valid mts value"))
	}
	request := CheckoutStart{
		Mts:           mts,
		OosArticleIds: nil,
	}
	return c.startCheckout(&request)
}

func (c *Client) CheckoutWithResolveKey(mts string, resolveKey string) (*Checkout, *CheckoutError) {
	if !c.IsAuthenticated() {
		return nil, wrapCheckoutError(authenticationError())
	}
	if strings.TrimSpace(mts) == "" || strings.TrimSpace(resolveKey) == "" {
		return nil, wrapCheckoutError(createError("CheckoutWithResolveKey requires a valid mts and resolveKey value"))
	}
	request := CheckoutStart{
		Mts:           mts,
		OosArticleIds: nil,
		ResolveKey:    "age_verified",
	}
	return c.startCheckout(&request)
}

func (c *Client) startCheckout(request *CheckoutStart) (*Checkout, *CheckoutError) {
	startCheckoutUrl := c.baseURL + "/cart/checkout/start"
	var checkout = Checkout{}
	err := c.post(startCheckoutUrl, request, &checkout)
	if err != nil {
		var re *CheckoutError
		ok := errors.As(err, &re)
		if ok {
			return nil, re
		} else {
			return nil, wrapCheckoutError(err)
		}
	}
	return &checkout, nil
}

func (c *Client) getCheckoutStatus(transactionId string) (string, error) {
	if !c.IsAuthenticated() {
		return "", wrapCheckoutError(authenticationError())
	}
	if strings.TrimSpace(transactionId) == "" {
		return "", wrapCheckoutError(createError("getCheckoutStatus requires a valid transactionId value"))
	}
	checkoutStatusUrl := c.baseURL + "/cart/checkout/" + transactionId + "/status"
	var checkoutStatus struct {
		CheckoutStatus string `json:"checkout_status"`
	}
	err := c.get(checkoutStatusUrl, &checkoutStatus)
	if err != nil {
		return "", err
	}
	return checkoutStatus.CheckoutStatus, nil
}

func (c *Client) CancelCheckout(transactionId string) error {
	if !c.IsAuthenticated() {
		return wrapCheckoutError(authenticationError())
	}
	if strings.TrimSpace(transactionId) == "" {
		return wrapCheckoutError(createError("CancelCheckout requires a valid transactionId value"))
	}
	cancelCheckoutUrl := c.baseURL + "/cart/checkout/cancel"
	var cancelTransactionRequest struct {
		TransactionId string `json:"transaction_id"`
	}
	cancelTransactionRequest.TransactionId = transactionId
	return c.post(cancelCheckoutUrl, cancelTransactionRequest, nil)
}
