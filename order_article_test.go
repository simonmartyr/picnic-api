package picnic

import "testing"

func TestOrderArticle_IsAvailable_False(t *testing.T) {
	oa := OrderArticle{
		Decorators: []Decorator{
			{
				Type: "UNAVAILABLE",
			},
		},
	}
	if oa.IsAvailable() {
		t.Error("Invalid Availability")
	}
}

func TestOrderArticle_IsAvailable_True(t *testing.T) {
	oa := OrderArticle{
		Decorators: []Decorator{},
	}
	if !oa.IsAvailable() {
		t.Error("Invalid Availability")
	}
}

func TestOrderArticle_Quantity(t *testing.T) {
	expected := 5
	oa := OrderArticle{
		Decorators: []Decorator{
			{
				Type:     "QUANTITY",
				Quantity: expected,
			},
		},
	}
	if oa.Quantity() != expected {
		t.Error("Invalid Quantity")
	}
}

func TestOrderArticle_Quantity_Default(t *testing.T) {
	oa := OrderArticle{
		Decorators: []Decorator{
			{
				Type: "UNAVAILABLE",
			},
		},
	}
	if oa.Quantity() != 0 {
		t.Error("Invalid Quantity")
	}
}
