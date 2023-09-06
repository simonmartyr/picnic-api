package picnic

import (
	"net/url"
	"strings"
)

type SearchSuggestion struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Links      []Link `json:"links"`
	Suggestion string `json:"suggestion"`
}

// GetSearchSuggestions Queries for terms that could be used for Search
// For example the query Pepper could result in 'ChiliPepper'
//
// Method requires client to be authenticated
func (c *Client) GetSearchSuggestions(query string) (*[]SearchSuggestion, error) {
	if !c.IsAuthenticated() {
		return nil, authenticationError()
	}
	if strings.TrimSpace(query) == "" {
		return nil, createError("GetSearchSuggestions requires a query string")
	}
	searchUrl := c.baseURL + "/suggest?search_term=" + url.QueryEscape(query)
	var searchSuggestion []SearchSuggestion
	err := c.get(searchUrl, &searchSuggestion)
	if err != nil {
		return nil, err
	}
	return &searchSuggestion, nil
}
