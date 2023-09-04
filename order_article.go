package picnic

type OrderArticle struct {
	Type             string      `json:"type"`
	Id               string      `json:"id"`
	Name             string      `json:"name"`
	DisplayPrice     int         `json:"display_price"`
	Price            int         `json:"price"`
	ImageId          string      `json:"image_id"`
	Images           []Image     `json:"images"`
	UnitQuantity     string      `json:"unit_quantity"`
	MaxOrderQuantity int         `json:"max_order_quantity"`
	Decorators       []Decorator `json:"decorators"`
}

func (a OrderArticle) Quantity() int {
	for _, decorator := range a.Decorators {
		if decorator.Type == "QUANTITY" {
			return decorator.Quantity
		}
	}
	return 0
}

func (a OrderArticle) IsAvailable() bool {
	for _, decorator := range a.Decorators {
		if decorator.Type == "UNAVAILABLE" {
			return false
		}
	}
	return true
}
