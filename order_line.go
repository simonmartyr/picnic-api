package picnic

type OrderLine struct {
	Type  string         `json:"type"`
	Id    string         `json:"id"`
	Items []OrderArticle `json:"items"`
}
