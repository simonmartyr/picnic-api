package picnic

import (
	"net/http"
	"testing"
)

func TestGetDeliveryScenario(t *testing.T) {
	deliveryId := "delivery500"
	c, s := testClientFile(http.StatusOK, "test/delivery_scenario.json")
	defer s.Close()
	res, err := c.GetDeliveryScenario(deliveryId)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Scenario) == 0 {
		t.Error("Scenario empty")
	}
	if res.Scenario[0].Lng != 4.9219651 {
		t.Error("Invalid lng")
	}
	if res.Destination.Street != "Cool Street" {
		t.Error("Invalid Destination Street")
	}
	if res.Driver.Name != "My Guy" {
		t.Error("Invalid Driver Name")
	}
}

func TestGetDeliveryScenario_Error_MissingId(t *testing.T) {
	deliveryId := "  "
	c, s := testClientFile(http.StatusOK, "test/delivery_scenario.json")
	defer s.Close()
	res, err := c.GetArticleDetails(deliveryId)
	if res != nil {
		t.Error("Invalid unexpected response")
	}
	if err == nil {
		t.Error("No error raised")
	}
}
