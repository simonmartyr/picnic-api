package picnic

type LoginInput struct {
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	ClientId int    `json:"client_id"`
}
