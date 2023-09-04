package picnic

type DeliverySlot struct {
	SlotId               string `json:"slot_id"`
	HubId                string `json:"hub_id"`
	FcId                 string `json:"fc_id"`
	WindowStart          string `json:"window_start"`
	CutOffTime           string `json:"cut_off_time"`
	IsAvailable          bool   `json:"is_available"`
	Icon                 Icon   `json:"icon"`
	Selected             bool   `json:"selected"`
	Reserved             bool   `json:"reserved"`
	MinimumOrderValue    int    `json:"minimum_order_value"`
	UnavailabilityReason string `json:"unavailability_reason"`
}
