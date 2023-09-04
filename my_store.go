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
