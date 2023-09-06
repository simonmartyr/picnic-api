package picnic

type MyStore struct {
	Type           string                   `json:"type"`
	Catalog        []Category               `json:"catalog"`
	Content        []DisplayPositionContent `json:"content"`
	User           User                     `json:"user"`
	FirstTimeUser  bool                     `json:"first_time_user"`
	LandingPageHit string                   `json:"landing_page_hit"`
	Id             string                   `json:"id"`
	Links          []Link                   `json:"links"`
}

// GetMyStore query to retrieve the MyStore for the current authenticated user
//
// Method requires client to be authenticated
func (c *Client) GetMyStore() (*MyStore, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	myStoreUrl := c.baseURL + "/my_store"
	var myStore MyStore
	err := c.get(myStoreUrl, &myStore)
	if err != nil {
		return nil, err
	}
	return &myStore, nil
}
