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
