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

func TestClient_ClearCart(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/cleared_cart.json")
	defer s.Close()
	res, err := c.ClearCart()
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Items) != 0 {
		t.Error("Invalid Item length")
	}
	if res.TotalPrice != 0 {
		t.Error("Invalid price")
	}
}

func TestClient_AddToCart(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/order_data.json")
	defer s.Close()
	res, err := c.AddToCart("s1005863", 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Items) == 0 {
		t.Error("Invalid Item length")
	}
}

func TestClient_RemoveFromCart(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/order_data.json")
	defer s.Close()
	res, err := c.RemoveFromCart("s1005863", 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Items) == 0 {
		t.Error("Invalid Item length")
	}
}

func TestClient_SetDeliverySlot(t *testing.T) {
	slotId := "64deb48e5e91961aea7ee0b0"
	c, s := testClientFile(http.StatusOK, "test/order_data.json")
	defer s.Close()
	res, err := c.SetDeliverySlot(slotId)
	if err != nil {
		t.Fatal(err)
	}
	if res.SelectedSlot.SlotId != slotId {
		t.Error("Invalid slotId")
	}
}
