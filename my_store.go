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

func (c *Client) GetMyStore() (*MyStore, error) {
	myStoreUrl := c.baseURL + "/my_store"
	var myStore MyStore
	err := c.get(myStoreUrl, &myStore)
	if err != nil {
		return nil, err
	}
	return &myStore, nil
}
