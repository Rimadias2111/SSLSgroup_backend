package swag

type CreateUpdatePerformance struct {
	Reason     string `json:"reason"`
	WhoseFault string `json:"whose_fault"`
	Status     string `json:"status"`
	Section    string `json:"section"`
	DisputedBy string `json:"disputed_by"`
	Company    string `json:"company"`
	LoadId     string `json:"load_id"`
}
