package picnic

import (
	"net/http"
	"testing"
)

func TestClient_InitiatePayment(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/initiate_payment_data.json")
	defer s.Close()
	res, err := c.InitiatePayment("id")
	if err != nil {
		t.Fatal(err)
	}
	if res.IssuerAuthenticationUrl != "https://averycoolurl.com" {
		t.Error("Invalid issuer auth url")
	}
	if res.Action.RedirectUrl != "https://averycoolurl.com" {
		t.Error("Invalid redirect")
	}
	if res.Action.Type != "REDIRECT" {
		t.Error("Invalid action type")
	}
}
