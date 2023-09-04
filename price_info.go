package picnic

type PriceInfo struct {
	Price         int    `json:"price"`
	PriceColor    string `json:"price_color"`
	OriginalPrice int    `json:"original_price"`
	Deposit       int    `json:"deposit"`
	BasePriceText string `json:"base_price_text"`
}
