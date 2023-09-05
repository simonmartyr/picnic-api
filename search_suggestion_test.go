package picnic

import (
	"net/http"
	"testing"
)

func Test_GetSuggestions(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/search_suggestion_data.json")
	defer s.Close()
	res, err := c.GetSearchSuggestions("pepper")
	if err != nil {
		t.Fatal(err)
	}
	results := *res
	if results[0].Suggestion != "pepper" {
		t.Error("Invalid Suggestion")
	}
	if len(results[0].Links) != 1 {
		t.Error("Invalid Links")
	}
}
