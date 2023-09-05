package picnic

import (
	"net/http"
	"testing"
)

func Test_GetMyStore(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/my_store_data.json")
	defer s.Close()
	res, err := c.GetMyStore()
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Catalog) == 0 {
		t.Error("Invalid Item Catalog length")
	}
}
