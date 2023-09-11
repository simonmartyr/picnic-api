package picnic

type OrderLine struct {
	Type         string         `json:"type"`
	Id           string         `json:"id"`
	Items        []OrderArticle `json:"items"`
	DisplayPrice int            `json:"display_price"`
	Price        int            `json:"price"`
	Decorators   []Decorator    `json:"decorators"`
}

func (o OrderLine) PriceIncludingPromotions() int {
	for _, d := range o.Decorators {
		if d.Type == "PRICE" {
			return d.DisplayPrice
		}
	}
	return o.DisplayPrice
}

func (s OrderLine) IsOnPromotion() bool {
	for _, d := range s.Decorators {
		if d.Type == "PROMO" {
			return true
		}
	}
	return false
}
