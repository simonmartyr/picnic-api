package picnic

import "net/url"

type SearchResult struct {
	Type  string          `json:"type"`
	Id    string          `json:"id"`
	Links []Link          `json:"links"`
	Name  string          `json:"name"`
	Items []SingleArticle `json:"items"`
	Level int             `json:"level"`
}

func (c *Client) Search(query string) (*[]SearchResult, error) {
	searchUrl := c.baseURL + "/search?search_term=" + url.QueryEscape(query)
	var searchResults []SearchResult
	err := c.get(searchUrl, &searchResults)
	if err != nil {
		return nil, err
	}
	return &searchResults, nil
}
