package picnic

type CheckoutError struct {
	Code       string
	Message    string
	ResolveKey string
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
