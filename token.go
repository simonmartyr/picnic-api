package picnic

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type token struct {
	Sub    string `json:"sub"`
	PcClid int    `json:"pc:clid"`
	PcDid  string `json:"pc:did"`
}

func (c *Client) parseJwt() error {
	tokenParts := strings.Split(c.token, ".")
	if len(tokenParts) != 3 {
		return errors.New("invalid token structure")
	}
	jsonContent, decodeErr := base64.RawStdEncoding.DecodeString(tokenParts[1])
	if decodeErr != nil {
		return decodeErr
	}
	var toReturn = token{}
	jsonErr := json.Unmarshal(jsonContent, &toReturn)
	if jsonErr != nil {
		return jsonErr
	}
	c.parsedToken = &toReturn
	return nil
}
