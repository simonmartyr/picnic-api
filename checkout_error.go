package picnic

import (
	"fmt"
)

type CheckoutError struct {
	Code       string
	Title      string
	Message    string
	ResolveKey string
	Blocking   bool
	Err        error
}

func (e *CheckoutError) Error() string {
	return e.Err.Error()
}

func wrapCheckoutError(err error) *CheckoutError {
	return &CheckoutError{
		Err: err,
	}
}

func (p *CheckoutError) IsAgeVerificationCheck() bool {
	return p.Code == "LEGACY_ALCOHOL_AGE_VERIFICATION_REQUIRED"
}

func (p *PicnicError) CreateCheckoutError() *CheckoutError {
	return &CheckoutError{
		Code:       p.Details.Type,
		Title:      p.Details.LocalizedTitle,
		Message:    p.Details.LocalizedMessage,
		ResolveKey: p.Details.ResolveKey,
		Blocking:   p.Details.Blocking,
		Err:        fmt.Errorf("picnic-api: cart has an issue %s with message %s", p.Code, p.Message),
	}
}
