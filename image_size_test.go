package picnic

import "testing"

func TestImageSize_String_Tiny(t *testing.T) {
	if Tiny.String() != "tiny" {
		t.Error("Tiny incorrect")
	}
}
func TestImageSize_String_Small(t *testing.T) {
	if Small.String() != "small" {
		t.Error("Small incorrect")
	}
}
func TestImageSize_String_Medium(t *testing.T) {
	if Medium.String() != "medium" {
		t.Error("Medium incorrect")
	}
}
func TestImageSize_String_Large(t *testing.T) {
	if Large.String() != "large" {
		t.Error("Large incorrect")
	}
}
func TestImageSize_String_ExtraLarge(t *testing.T) {
	if ExtraLarge.String() != "extra-large" {
		t.Error("ExtraLarge incorrect")
	}
}
func TestImageSize_String_Default(t *testing.T) {
	var c ImageSize = 9
	if c.String() != "small" {
		t.Error("Default incorrect")
	}
}
