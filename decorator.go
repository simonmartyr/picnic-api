package picnic

type Decorator struct {
	Type              string      `json:"type"`
	HeightPercentage  float64     `json:"height_percentage,omitempty"`
	ImageIDs          []string    `json:"image_ids,omitempty"`
	Banners           []SubBanner `json:"banners,omitempty"`
	BasePriceText     string      `json:"base_price_text,omitempty"`
	Period            string      `json:"period,omitempty"`
	Label             string      `json:"label,omitempty"`
	Link              Link        `json:"link,omitempty"`
	Images            []string    `json:"images,omitempty"`
	SellableItemCount int         `json:"sellable_item_count,omitempty"`
	DisplayPrice      int         `json:"display_price,omitempty"`
	Quantity          int         `json:"quantity,omitempty"`
	Styles            []Style     `json:"styles,omitempty"`
	UnitQuantityText  string      `json:"unit_quantity_text,omitempty"`
	ValidUntil        string      `json:"valid_until,omitempty"`
	Reason            string      `json:"reason,omitempty"`
	//Explanation       string      `json:"explanation,omitempty"`
}
