package picnic

type ReturnContainer struct {
	Type          string `json:"type"`
	LocalizedName string `json:"localized_name"`
	Quantity      int    `json:"quantity"`
	Price         int    `json:"price"`
}
