package picnic

type HouseholdDetails struct {
	Adults   int    `json:"adults"`
	Children int    `json:"children"`
	Cats     int    `json:"cats"`
	Dogs     int    `json:"dogs"`
	Author   string `json:"author"`
}
