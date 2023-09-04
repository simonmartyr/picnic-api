package picnic

type ArticleDetails struct {
	Type             string             `json:"type"`
	Id               string             `json:"id"`
	Name             string             `json:"name"`
	PriceInfo        PriceInfo          `json:"price_info"`
	Images           []Image            `json:"images"`
	UnitQuantity     string             `json:"unit_quantity"`
	MaxOrderQuantity int                `json:"max_order_quantity"`
	Decorators       []Decorator        `json:"decorators"`
	Description      ArticleDescription `json:"description"`
}

func (c *Client) GetArticleDetails(articleId string) (*ArticleDetails, error) {
	articleDetailsUrl := c.baseURL + "/articles/" + articleId
	var article ArticleDetails
	err := c.get(articleDetailsUrl, &article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}
