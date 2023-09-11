package picnic

type SingleArticle struct {
	Type         string      `json:"type"`
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	DisplayPrice int         `json:"display_price"`
	Price        int         `json:"price"`
	ImageId      string      `json:"image_id"`
	UnitQuantity string      `json:"unit_quantity"`
	Decorators   []Decorator `json:"decorators"`
}

func (s SingleArticle) IsOnPromotion() bool {
	for _, d := range s.Decorators {
		if d.Type == "PROMO" {
			return true
		}
	}
	return false
}

func (o SingleArticle) PriceIncludingPromotions() int {
	for _, d := range o.Decorators {
		if d.Type == "PRICE" {
			return d.DisplayPrice
		}
	}
	return o.DisplayPrice
}
