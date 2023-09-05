package picnic

import (
	"net/http"
	"testing"
)

func TestGetArticleDetails(t *testing.T) {
	articleId := "s1005863"
	c, s := testClientFile(http.StatusOK, "test/article_details_data.json")
	defer s.Close()
	res, err := c.GetArticleDetails(articleId)
	if err != nil {
		t.Fatal(err)
	}
	if res.Name != "Koffiecups recyclezak" {
		t.Error("Invalid article name")
	}
	if res.Description.Main == "" {
		t.Error("Invalid description")
	}
	if len(res.Images) != 1 {
		t.Error("Invalid images")
	}
}
