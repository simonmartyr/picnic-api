package picnic

import (
	"errors"
	"fmt"
)

const (
	baseError             = "picnic-api:"
	missingAuthentication = baseError + " Method requires authentication, please login"
)

func createError(message string) error {
	return fmt.Errorf("%s %s", baseError, message)
}

func authenticationError() error {
	return errors.New(missingAuthentication)
}
