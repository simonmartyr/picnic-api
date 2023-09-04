package picnic

type DisplayPositionContent struct {
	Type            string    `json:"type"`
	Id              string    `json:"id"`
	Links           Link      `json:"links"`
	DisplayPosition string    `json:"display_position"`
	items           []Content `json:"items"`
}
