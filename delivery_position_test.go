package picnic

import (
	"net/http"
	"testing"
)

func TestGetDeliveryPosition(t *testing.T) {
	deliveryId := "delivery500"
	c, s := testClientFile(http.StatusOK, "test/delivery_position.json")
	defer s.Close()
	res, err := c.GetDeliveryPosition(deliveryId)
	if err != nil {
		t.Fatal(err)
	}
	if res.ScenarioTs != 1704795199346 {
		t.Error("Invalid ts")
	}
	if res.Eta != 1704797648363 {
		t.Error("Invalid eta")
	}
	if res.QueryInterval != 10000 {
		t.Error("Invalid interval")
	}
	if res.EtaWindow.Start != "2024-01-09T11:43:26.361+01:00" {
		t.Error("Invalid Eta Window Start")
	}
	if res.EtaWindow.End != "2024-01-09T12:03:26.361+01:00" {
		t.Error("Invalid Eta Window End")
	}
}

func TestGetDeliveryPosition_Error_MissingId(t *testing.T) {
	deliveryId := "  "
	c, s := testClientFile(http.StatusOK, "test/delivery_position.json")
	defer s.Close()
	res, err := c.GetArticleDetails(deliveryId)
	if res != nil {
		t.Error("Invalid unexpected response")
	}
	if err == nil {
		t.Error("No error raised")
	}
}
