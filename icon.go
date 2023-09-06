package picnic

type Icon struct {
	PmlVersion         string             `json:"pml_version"`
	Component          Component          `json:"component"`
	Images             Image              `json:"images"`
	TrackingAttributes TrackingAttributes `json:"tracking_attributes"`
}
