package picnic

type Decorator struct {
	Type              string      `json:"type"`
	HeightPercentage  float64     `json:"height_percentage"`
	ImageIDs          []string    `json:"image_ids"`
	Banners           []SubBanner `json:"banners"`
	BasePriceText     string      `json:"base_price_text"`
	Period            string      `json:"period"`
	Label             string      `json:"label"`
	Link              Link        `json:"link"`
	Images            []string    `json:"images"`
	SellableItemCount int         `json:"sellable_item_count"`
	DisplayPrice      int         `json:"display_price"`
	Quantity          int         `json:"quantity"`
	Styles            []Style     `json:"styles"`
	UnitQuantityText  string      `json:"unit_quantity_text"`
	ValidUntil        string      `json:"valid_until"`
	Reason            string      `json:"reason"`
	Text              string      `json:"text"`
}
