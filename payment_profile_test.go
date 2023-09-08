package picnic

import (
	"net/http"
	"testing"
)

func TestClient_GetPaymentProfile(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/payment_profile_data.json")
	defer s.Close()
	res, err := c.GetPaymentProfile()
	if err != nil {
		t.Fatal(err)
	}
	if len(res.StoredPaymentOptions) != 5 {
		t.Error("Invalid Payment Options")
	}
	if res.StoredPaymentOptions[0].Brand != "BrandName" {
		t.Error("Invalid Payment Options brand")
	}
	if len(res.AvailablePaymentMethods) != 2 {
		t.Error("Invalid Payment Methods")
	}
	if len(res.AvailablePaymentMethods) != 2 {
		t.Error("Invalid Available Payment Options")
	}
	if len(res.PaymentMethods) != 2 {
		t.Error("Invalid Payment methods")
	}
	if res.PreferredPaymentOptionId != "de1f471e-a461-4dc2-a2fc-a618ddc9b6c4" {
		t.Error("Invalid preferred payment")
	}
}

func TestClient_GetPaymentProfile_RequiresAuth(t *testing.T) {
	c := &Client{
		http:  http.DefaultClient,
		token: "",
	}
	res, err := c.GetPaymentProfile()
	if res != nil {
		t.Error("Unexpected response")
	}
	if err.Error() != authenticationError().Error() {
		t.Error("Incorrect error")
	}
}
