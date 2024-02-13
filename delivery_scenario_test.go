package picnic

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"
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

func Test_Integration_Deliveries(t *testing.T) {
	godotenv.Load()
	if os.Getenv("SKIP_WIP") != "" {
		//TODO https://github.com/simonmartyr/picnic-api/issues/4
		t.Skip("Skipping")
	}
	c := New(&http.Client{},
		WithUsername(os.Getenv("USERNAME")),
		WithHashedPassword(os.Getenv("SECRET")),
		WithVersion("15"),
	)
	authErr := c.Authenticate()
	if authErr != nil {
		t.Error("auth failed")
	}
	deliveries, dErr := c.GetDeliveries([]DeliveryStatus{COMPLETED, CURRENT})
	if dErr != nil {
		t.Error(dErr)
	}
	pos, posErr := c.GetDeliveryPosition((*deliveries)[0].DeliveryId)
	if posErr != nil {
		t.Fatal(posErr)
	}
	if pos == nil {
		t.Error("invalid pos")
	}
}
