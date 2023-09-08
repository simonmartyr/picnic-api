package picnic

type PicnicError struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Details PicnicErrorDetails `json:"details"`
}

type PicnicErrorDetails struct {
	Type                        string `json:"type"`
	LocalizedMessage            string `json:"localized_message"`
	Blocking                    bool   `json:"blocking"`
	ResolveKey                  string `json:"resolve_key"`
	LocalizedTitle              string `json:"localized_title"`
	LocalizedContinueButtonText string `json:"localized_continue_button_text"`
	LocalizedCancelButtonText   string `json:"localized_cancel_button_text"`
}

func (p PicnicError) IsAgeVerificationCheck() bool {
	return p.Details.Type == "LEGACY_ALCOHOL_AGE_VERIFICATION_REQUIRED"
}
