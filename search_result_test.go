package picnic

import (
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/search_data.json")
	defer s.Close()
	res, err := c.Search("pepper")
	if err != nil {
		t.Fatal(err)
	}
	results := *res
	if len(results) != 1 {
		t.Error("Invalid result length")
	}
	if results[0].Items[0].Name != "Dr Pepper regular" {
		t.Error("Invalid item name")
	}
}

func TestSearch_Error(t *testing.T) {
	c, s := testClientFile(http.StatusUnauthorized, "test/error.json")
	defer s.Close()
	res, err := c.Search("pepper")
	if res != nil {
		t.Error("Unexpected result")
	}
	if err == nil {
		t.Error("No error produced")
	}
}

func TestSearch_RequiresAuth(t *testing.T) {
	c := &Client{
		http:  http.DefaultClient,
		token: "",
	}
	res, err := c.Search("query")
	if res != nil {
		t.Error("Unexpected response")
	}
	if err.Error() != authenticationError().Error() {
		t.Error("Incorrect error")
	}
}

func TestSearch_RequiresTerm(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/search_data.json")
	defer s.Close()
	res, err := c.Search(" ")
	if res != nil {
		t.Error("Unexpected response")
	}
	if err == nil {
		t.Error("Error missing")
	}
}
