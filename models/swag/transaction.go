package swag

type CreateUpdateTransaction struct {
	From         string  `json:"from"`
	To           string  `json:"to"`
	PuTime       string  `json:"pu_time"`
	DeliveryTime string  `json:"delivery_time"`
	Success      bool    `json:"success"`
	LoadedMiles  int64   `json:"loaded_miles"`
	TotalMiles   int64   `json:"total_miles"`
	Provider     string  `json:"provider"`
	Cost         int64   `json:"cost"`
	Rate         float64 `json:"rate"`
	DriverId     string  `json:"driver_id"`
	EmployeeId   string  `json:"employee_id"`
	CargoID      string  `json:"cargo_id"`
}
