package picnic

type User struct {
	Id                  string           `json:"user_id"`
	Firstname           string           `json:"firstname"`
	Lastname            string           `json:"lastname"`
	Address             Address          `json:"address"`
	Phone               string           `json:"phone"`
	ContactEmail        string           `json:"contact_email"`
	FeatureToggles      []interface{}    `json:"feature_toggles"`
	PushSubscriptions   []Subscription   `json:"push_subscriptions"`
	Subscriptions       []Subscription   `json:"subscriptions"`
	CustomerType        string           `json:"customer_type"`
	HouseholdDetails    HouseholdDetails `json:"household_details"`
	CheckGeneralConsent bool             `json:"check_general_consent"`
	PlacedOrder         bool             `json:"placed_order"`
	ReceivedDelivery    bool             `json:"received_delivery"`
	TotalDeliveries     int              `json:"total_deliveries"`
	CompletedDeliveries int              `json:"completed_deliveries"`
	ConsentDecisions    ConsentDecisions `json:"consent_decisions"`
}

// GetUser Retrieves the details of currently authenticated user.
//
// Method requires client to be authenticated
func (c *Client) GetUser() (*User, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	userUrl := c.baseURL + "/user"
	var user User
	err := c.get(userUrl, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
