package picnic

import (
	"net/http"
	"testing"
)

func TestGetCart(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/order_data.json")
	defer s.Close()
	res, err := c.GetCart()
	if err != nil {
		t.Fatal(err)
	}
	if res.Items[0].Items[0].Id != "s1005863" {
		t.Error("Invalid Item Id")
	}
	if res.Items[0].Items[0].Name != "Koffiecups recyclezak" {
		t.Error("Invalid Item Name")
	}
}
