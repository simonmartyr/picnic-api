package picnic

import (
	"net/http"
	"testing"
)

func Test_GetDeliveries(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/deliveries_data.json")
	defer s.Close()
	var filter = []DeliveryStatus{COMPLETED}
	res, err := c.GetDeliveries(filter)
	if err != nil {
		t.Fatal(err)
	}
	results := *res
	if results[0].Slot.SlotId != "64dac0105e91961aea7ed87c" {
		t.Error("Invalid DeliverySlotId")
	}
	if results[0].Status != COMPLETED {
		t.Error("Invalid Delivery Status")
	}
}

func Test_GetDeliveries_nil(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/deliveries_data.json")
	defer s.Close()
	res, err := c.GetDeliveries(nil)
	if err != nil {
		t.Fatal(err)
	}
	results := *res
	if results[0].Slot.SlotId != "64dac0105e91961aea7ed87c" {
		t.Error("Invalid DeliverySlotId")
	}
	if results[0].Status != COMPLETED {
		t.Error("Invalid Delivery Status")
	}
}

func Test_GetDeliveries_Error_RequiresAuth(t *testing.T) {
	var filter = []DeliveryStatus{COMPLETED}
	c := &Client{
		http:  http.DefaultClient,
		token: "",
	}
	res, err := c.GetDeliveries(filter)
	if res != nil {
		t.Error("Unexpected response")
	}
	if err.Error() != authenticationError().Error() {
		t.Error("Incorrect error")
	}
}

func Test_GetDeliveryById(t *testing.T) {
	id := "a3szungyku"
	c, s := testClientFile(http.StatusOK, "test/delivery_data.json")
	defer s.Close()
	res, err := c.GetDelivery(id)
	if err != nil {
		t.Fatal(err)
	}
	if res.Slot.SlotId != "64bc6d892463c8138fddfe95" {
		t.Error("Invalid DeliverySlotId")
	}
	if res.Status != COMPLETED {
		t.Error("Invalid Delivery Status")
	}
	if res.Id != id {
		t.Error("Invalid Delivery id")
	}
}

func Test_GetDeliveryById_Errors_Bad_Id(t *testing.T) {
	id := "  "
	c, s := testClientFile(http.StatusOK, "test/delivery_data.json")
	defer s.Close()
	res, err := c.GetDelivery(id)
	if res != nil {
		t.Error("Invalid unexpected response")
	}
	if err == nil {
		t.Error("No error raised")
	}
}

func Test_GetDeliveryById_Errors_RequiresAuth(t *testing.T) {
	id := "a3szungyku"
	c := &Client{
		http:  http.DefaultClient,
		token: "",
	}
	res, err := c.GetDelivery(id)
	if res != nil {
		t.Error("Unexpected response")
	}
	if err.Error() != authenticationError().Error() {
		t.Error("Incorrect error")
	}
}
