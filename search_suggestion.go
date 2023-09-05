package picnic

import "net/url"

type SearchSuggestion struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Links      []Link `json:"links"`
	Suggestion string `json:"suggestion"`
}

func (c *Client) GetSearchSuggestions(query string) (*[]SearchSuggestion, error) {
	searchUrl := c.baseURL + "/suggest?search_term=" + url.QueryEscape(query)
	var searchSuggestion []SearchSuggestion
	err := c.get(searchUrl, &searchSuggestion)
	if err != nil {
		return nil, err
	}
	return &searchSuggestion, nil
}
