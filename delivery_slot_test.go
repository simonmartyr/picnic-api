package picnic

import (
	"net/http"
	"testing"
)

func Test_GetDeliverySlots(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/delivery_slots_data.json")
	defer s.Close()
	res, err := c.GetDeliverySlots()
	if err != nil {
		t.Fatal(err)
	}
	if res.DeliverySlots[0].SlotId != "64e3fa8ad01b0f2489ebfcdf" {
		t.Error("Invalid DeliverySlotId")
	}
	if res.SelectedSlot.SlotId != "64e69d9e0a1c734a4e1b285c" {
		t.Error("Invalid Selected Slot")
	}
}

func Test_GetDeliverySlots_Errors_RequiresAuth(t *testing.T) {
	c := &Client{
		http:  http.DefaultClient,
		token: "",
	}
	res, err := c.GetDeliverySlots()
	if res != nil {
		t.Error("Unexpected response")
	}
	if err.Error() != authenticationError().Error() {
		t.Error("Incorrect error")
	}
}
