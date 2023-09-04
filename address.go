package picnic

type Address struct {
	HouseNumber    int    `json:"house_number"`
	HouseNumberExt string `json:"house_number_ext"`
	Postcode       string `json:"postcode"`
	Street         string `json:"street"`
	City           string `json:"city"`
}
