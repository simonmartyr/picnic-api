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

func TestGetArticleDetails_Error_MissingId(t *testing.T) {
	articleId := "  "
	c, s := testClientFile(http.StatusOK, "test/article_details_data.json")
	defer s.Close()
	res, err := c.GetArticleDetails(articleId)
	if res != nil {
		t.Error("Invalid unexpected response")
	}
	if err == nil {
		t.Error("No error raised")
	}
}

func TestGetArticleDetails_Error_RequiresAuth(t *testing.T) {
	articleId := "s1005863"
	c := &Client{
		http:  http.DefaultClient,
		token: "",
	}
	res, err := c.GetArticleDetails(articleId)
	if res != nil {
		t.Error("Unexpected response")
	}
	if err.Error() != authenticationError().Error() {
		t.Error("Incorrect error")
	}
}
