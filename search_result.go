package picnic

import (
	"net/url"
	"strings"
)

type SearchResult struct {
	Type  string          `json:"type"`
	Id    string          `json:"id"`
	Links []Link          `json:"links"`
	Name  string          `json:"name"`
	Items []SingleArticle `json:"items"`
	Level int             `json:"level"`
}

// Search Retrieves articles that relate to a given query. The results given are SingleArticle
// These have basic information about an article. To get more information about an article leverage GetArticleDetails.
//
// Method requires client to be authenticated
func (c *Client) Search(query string) (*[]SearchResult, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if strings.TrimSpace(query) == "" {
		return nil, createError("Search requires a query string")
	}
	searchUrl := c.baseURL + "/search?search_term=" + url.QueryEscape(query)
	var searchResults []SearchResult
	err := c.get(searchUrl, &searchResults)
	if err != nil {
		return nil, err
	}
	return &searchResults, nil
}
