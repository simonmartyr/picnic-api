package picnic

type LoginResult struct {
	UserId                             string `json:"user_id"`
	SecondFactorAuthenticationRequired string `json:"second_factor_authentication_required"`
	AuthKey                            string `json:"auth_key"`
}
