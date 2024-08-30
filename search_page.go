package picnic

import (
	"net/url"
	"strings"
)

type SearchPage struct {
	Body FirstChild `json:"body"`
}

type FirstChild struct {
	Child RecursiveChildren `json:"child"`
}

type RecursiveChildren struct {
	Children []RecursiveChildren `json:"children"`
	Content  SearchContent       `json:"content"`
}

type SearchContent struct {
	Type        string        `json:"type"`
	SellingUnit SingleArticle `json:"sellingUnit"`
}

const sellingUnitType = "SELLING_UNIT_TILE"

// SearchArticles Retrieves articles that relate to a given query. The results given are SingleArticle
// These have basic information about an article. To get more information about an article leverage GetArticleDetails.
//
// Method requires client to be authenticated
func (c *Client) SearchArticles(query string) ([]SingleArticle, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if strings.TrimSpace(query) == "" {
		return nil, createError("Search requires a query string")
	}
	searchUrl := c.baseURL + "/pages/search-page-results?search_term=" + url.QueryEscape(query)
	var searchPage SearchPage
	err := c.get(searchUrl, &searchPage)
	if err != nil {
		return nil, err
	}
	return searchPage.extractArticles(), nil
}

func (page *SearchPage) extractArticles() []SingleArticle {
	var articles []SingleArticle
	for _, child := range page.Body.Child.Children {
		child.findArticles(&articles)
	}
	return articles
}

func (recursive *RecursiveChildren) findArticles(articles *[]SingleArticle) {
	if &recursive.Content == nil && (recursive.Children == nil || len(recursive.Children) == 0) {
		return
	}
	if &recursive.Content != nil && recursive.Content.Type == sellingUnitType {
		*articles = append(*articles, recursive.Content.SellingUnit)
	}
	for _, child := range recursive.Children {
		child.findArticles(articles)
	}
}
