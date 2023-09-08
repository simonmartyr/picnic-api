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
	Mts           int    `json:"mts"`
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

// StartCheckout begin the process of payment for the current order.
// Upon success the Checkout result will contain an OrderId which can be used to initiate payment
// Using the method InitiatePayment.
// A CheckoutError occurs when articles or issues exist with the current order. In the example of orders which contain
// Alcohol, a call to CheckoutWithResolveKey is required, the resolveKey can be found within the CheckoutError
//
// Method requires client to be authenticated
func (c *Client) StartCheckout(mts int) (*Checkout, *CheckoutError) {
	if !c.IsAuthenticated() {
		return nil, wrapCheckoutError(authenticationError())
	}
	request := CheckoutStart{
		Mts:           mts,
		OosArticleIds: nil,
	}
	return c.startCheckout(&request)
}

// CheckoutWithResolveKey The same as StartCheckout but if an order is flagged with issues, leverage this method in order to set the required resolveKeys
//
// Method requires client to be authenticated
func (c *Client) CheckoutWithResolveKey(mts int, resolveKey string) (*Checkout, *CheckoutError) {
	if !c.IsAuthenticated() {
		return nil, wrapCheckoutError(authenticationError())
	}
	if strings.TrimSpace(resolveKey) == "" {
		return nil, wrapCheckoutError(createError("CheckoutWithResolveKey requires a valid resolveKey value"))
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

// GetCheckoutStatus Whilst a checkout process is ongoing, calls to this endpoint will report its current status
// This can be helpful to verify if the payment has been received
//
// Method requires client to be authenticated
func (c *Client) GetCheckoutStatus(transactionId string) (string, error) {
	if !c.IsAuthenticated() {
		return "", wrapCheckoutError(authenticationError())
	}
	if strings.TrimSpace(transactionId) == "" {
		return "", wrapCheckoutError(createError("GetCheckoutStatus requires a valid transactionId value"))
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

// CancelCheckout ends the transaction with the associated id
//
// Method requires client to be authenticated
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
