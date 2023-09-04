package picnic

type Subscription struct {
	ListId     string `json:"list_id"`
	Subscribed bool   `json:"subscribed"`
	Name       string `json:"name"`
}
