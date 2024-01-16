package picnic

import "testing"

func TestParseJwt(t *testing.T) {
	client := &Client{
		token: mockToken,
	}
	tokenErr := client.parseJwt()
	if tokenErr != nil {
		t.Fatal(tokenErr)
	}
	if client.parsedToken.PcDid == "" {
		t.Error("PcDid invalid")
	}
}

func TestParseJwtInvalid(t *testing.T) {
	client := &Client{
		token: "fakeboi",
	}
	tokenErr := client.parseJwt()
	if tokenErr == nil {
		t.Error("token should not parse")
	}
}
