package picnic

type Content struct {
	Type            string           `json:"type"`
	Id              string           `json:"id"`
	DisplayPosition string           `json:"display_position"`
	Payload         TemplatedContent `json:"payload"`
}
