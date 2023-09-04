package picnic

type Category struct {
	Type                     string      `json:"type"`
	Id                       string      `json:"id"`
	Decorators               []Decorator `json:"decorators"`
	Links                    []Link      `json:"links"`
	Name                     string      `json:"name"`
	Items                    []Category  `json:"items"`
	level                    int         `json:"level"`
	IsIncludedInCategoryTree bool        `json:"is_included_in_category_tree"`
	hidden                   bool        `json:"hidden"`
}
