package picnic

type PaymentProfile struct {
	StoredPaymentOptions       []StoredPaymentOption    `json:"stored_payment_options"`
	AvailablePaymentMethods    []AvailablePaymentMethod `json:"available_payment_methods"`
	PaymentMethods             []PaymentMethod          `json:"payment_methods"`
	PreferredPaymentOptionId   string                   `json:"preferred_payment_option_id"`
	AvailablePaymentMethodItem string                   `json:"available_payment_method_item"`
}

type StoredPaymentOption struct {
	Id            string `json:"id"`
	PaymentMethod string `json:"payment_method"`
	Brand         string `json:"brand"`
	Account       string `json:"account"`
	DisplayName   string `json:"display_name"`
	IconUrl       string `json:"icon_url"`
}

type AvailablePaymentMethod struct {
	PaymentMethod  string `json:"payment_method"`
	AvailableBanks []Bank `json:"available_banks"`
}

type Bank struct {
	BankId string `json:"bank_id"`
	Name   string `json:"name"`
}

type PaymentMethod struct {
	PaymentMethod string         `json:"payment_method"`
	DisplayName   string         `json:"display_name"`
	IconUrl       string         `json:"icon_url"`
	Brands        []PaymentBrand `json:"brands"`
	Visibility    string         `json:"visibility"`
}

type PaymentBrand struct {
	Brand       string `json:"brand"`
	DisplayName string `json:"display_name"`
	IconUrl     string `json:"icon_url"`
}

// GetPaymentProfile retrieves the payment options for the authenticated user
//
// Method requires client to be authenticated
func (c *Client) GetPaymentProfile() (*PaymentProfile, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	paymentProfileUrl := c.baseURL + "/payment-profile"
	var profile PaymentProfile
	err := c.get(paymentProfileUrl, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
