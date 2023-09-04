package picnic

type SubBanner struct {
	BannerID    string            `json:"banner_id"`
	ImageID     string            `json:"image_id"`
	DisplayTime string            `json:"display_time"`
	Description string            `json:"description"`
	Reference   DeeplinkReference `json:"reference"`
	Position    string            `json:"position"`
}
