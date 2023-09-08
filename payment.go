package picnic

import "strings"

type InitiatePaymentRequest struct {
	AppReturnUrl string `json:"app_return_url"`
	OrderId      string `json:"order_id"`
}

type Payment struct {
	PaymentId               string        `json:"payment_id"`
	TransactionId           string        `json:"transaction_id"`
	IssuerAuthenticationUrl string        `json:"issuer_authentication_url"`
	Action                  PaymentAction `json:"action"`
}

type PaymentAction struct {
	Type        string `json:"type"`
	RedirectUrl string `json:"redirect_url"`
}

// InitiatePayment begin the process of paying for the order with a given orderId
// After payment is successfully started, GetCheckoutStatus can be used to monitor the status
//
// Method requires client to be authenticated
func (c *Client) InitiatePayment(orderId string) (*Payment, error) {
	if !c.IsAuthenticated() {
		return nil, wrapCheckoutError(authenticationError())
	}
	if strings.TrimSpace(orderId) == "" {
		return nil, wrapCheckoutError(createError("GetCheckoutStatus requires a valid transactionId value"))
	}
	initiatePaymentUrl := c.baseURL + "/cart/checkout/initiate_payment"
	var initiateRequest = InitiatePaymentRequest{
		AppReturnUrl: "nl.picnic-supermarkt://payment", //not sure if needed
		OrderId:      orderId,
	}
	var response = Payment{}
	err := c.post(initiatePaymentUrl, initiateRequest, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
