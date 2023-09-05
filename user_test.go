package picnic

import (
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/user_data.json")
	defer s.Close()
	res, err := c.GetUser()
	if err != nil {
		t.Fatal(err)
	}
	if res.Firstname != "Simon" {
		t.Error("Invalid User firstName")
	}
	if res.Lastname != "Martyr" {
		t.Error("Invalid User lastName")
	}
	if res.HouseholdDetails.Adults != 1 {
		t.Error("Invalid household adult count")
	}
	if res.Address.Street != "Cool Street" {
		t.Error("Invalid address")
	}
	if !res.ConsentDecisions.NespressoDataSharing {
		t.Error("Invalid nespresso consent")
	}
}

func TestGetUserError(t *testing.T) {
	c, s := testClientFile(http.StatusUnauthorized, "test/error.json")
	defer s.Close()
	_, err := c.GetUser()
	if err == nil {
		t.Error("err was nil")
	}
}
