package picnic

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"testing"
)

func TestSearchPage(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/search-page.json")
	defer s.Close()
	res, err := c.SearchArticles("melk")
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 79 {
		t.Error("Invalid result length")
	}
	for _, item := range res {
		if item.Name == "" {
			t.Error("Invalid Item Name")
		}
		if item.ImageId == "" {
			t.Error("Invalid Item Image Id")
		}
		if item.DisplayPrice == 0 {
			t.Error("Invalid Display Price")
		}
	}
}

func TestSearchPage_Error(t *testing.T) {
	c, s := testClientFile(http.StatusUnauthorized, "test/error.json")
	defer s.Close()
	res, err := c.SearchArticles("pepper")
	if res != nil {
		t.Error("Unexpected result")
	}
	if err == nil {
		t.Error("No error produced")
	}
}

func TestSearchPage_RequiresAuth(t *testing.T) {
	c := &Client{
		http:  http.DefaultClient,
		token: "",
	}
	res, err := c.SearchArticles("query")
	if res != nil {
		t.Error("Unexpected response")
	}
	if err.Error() != authenticationError().Error() {
		t.Error("Incorrect error")
	}
}

func TestSearchPage_RequiresTerm(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/search-page.json")
	defer s.Close()
	res, err := c.Search(" ")
	if res != nil {
		t.Error("Unexpected response")
	}
	if err == nil {
		t.Error("Error missing")
	}
}

func Test_Integration_SearchPage(t *testing.T) {
	godotenv.Load()
	c := New(&http.Client{},
		WithUsername(os.Getenv("USERNAME")),
		WithHashedPassword(os.Getenv("SECRET")),
		WithVersion("15"),
	)
	authErr := c.Authenticate()
	if authErr != nil {
		t.Error("auth failed")
	}
	res, err := c.SearchArticles("melk")
	if err != nil {
		t.Fatal(err)
	}
	if len(res) == 0 {
		t.Error("no results found")
	}
}
