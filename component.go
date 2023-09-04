package picnic

type Component struct {
	Type       string `json:"type"`
	Source     Source `json:"source"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	ResizeMode string `json:"resize_mode"`
}

type Source struct {
	Id string `json:"id"`
}
