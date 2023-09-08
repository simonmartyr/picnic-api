package picnic

import (
	"net/http"
	"testing"
)

func TestClient_StartCheckout(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/checkout_start_data.json")
	defer s.Close()
	res, err := c.StartCheckout(1)
	if err != nil {
		t.Fatal(*err)
	}
	if res.OrderId != "801-620-6076" {
		t.Error("Invalid Item length")
	}
	if res.TotalPrice != 12009 {
		t.Error("Invalid price")
	}
}

func TestClient_CheckoutWithResolveKey(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/checkout_start_data.json")
	defer s.Close()
	res, err := c.CheckoutWithResolveKey(1, "key")
	if err != nil {
		t.Fatal(err)
	}
	if res.OrderId != "801-620-6076" {
		t.Error("Invalid Item length")
	}
	if res.TotalPrice != 12009 {
		t.Error("Invalid price")
	}
}

func TestClient_StartCheckout_AgeError(t *testing.T) {
	c, s := testClientFile(http.StatusBadRequest, "test/age_verification_data.json")
	defer s.Close()
	res, err := c.StartCheckout(1)
	if res != nil {
		t.Error("Unexpected response")
	}
	if err.ResolveKey != "age_verified" {
		t.Error("Invalid resolve key")
	}
}

func TestClient_CancelCheckout(t *testing.T) {
	c, s := testClientFile(http.StatusNoContent, "test/empty.json")
	defer s.Close()
	err := c.CancelCheckout("id")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetCheckoutStatus(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/checkout_status_data.json")
	defer s.Close()
	res, err := c.GetCheckoutStatus("id")
	if err != nil {
		t.Fatal(err)
	}
	if res == "" {
		t.Error("Invalid status")
	}
	if res != "UNKNOWN" {
		t.Error("Unexpected status")
	}
}
