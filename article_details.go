package picnic

import (
	"strings"
)

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

// GetArticleDetails Retrieve a single article by its identifier. Article details provides additional information
// not found on the SingleArticle such as additional images and description.
//
// Method requires client to be authenticated
func (c *Client) GetArticleDetails(articleId string) (*ArticleDetails, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if strings.TrimSpace(articleId) == "" {
		return nil, createError("GetArticleDetails requires a valid articleId string")
	}
	articleDetailsUrl := c.baseURL + "/articles/" + articleId
	var article ArticleDetails
	err := c.get(articleDetailsUrl, &article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}
