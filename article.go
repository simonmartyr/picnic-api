package picnic

type Article struct {
	Type             string             `json:"type"`
	Id               string             `json:"id"`
	Name             string             `json:"name"`
	DisplayPrice     int                `json:"display_price"`
	PriceInfo        PriceInfo          `json:"price_info"`
	Price            int                `json:"price"`
	ImageId          string             `json:"image_id"`
	Images           []Image            `json:"images"`
	UnitQuantity     string             `json:"unit_quantity"`
	MaxOrderQuantity int                `json:"max_order_quantity"`
	Decorators       []Decorator        `json:"decorators"`
	Description      ArticleDescription `json:"description"`
}
