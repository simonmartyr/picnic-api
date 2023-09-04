package picnic

type ImageSize int64

const (
	Tiny ImageSize = iota
	Small
	Medium
	Large
	ExtraLarge
)

func (s ImageSize) String() string {
	switch s {
	case Tiny:
		return "tiny"
	case Small:
		return "small"
	case Medium:
		return "medium"
	case Large:
		return "large"
	case ExtraLarge:
		return "extra-large"
	}
	return "small"
}
