package picnic

type DeliveryStatus string

const (
	CANCELLED DeliveryStatus = "CANCELLED"
	COMPLETED DeliveryStatus = "COMPLETED"
	CURRENT   DeliveryStatus = "CURRENT"
)
