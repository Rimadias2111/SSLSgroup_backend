package swag

type CreateUpdateLogistic struct {
	DriverId string `json:"driver_id"`
	CargoId  string `json:"cargo_id"`
	Status   string `json:"status"`
	StTime   string `json:"st_time"`
	Location string `json:"location"`
	Notion   string `json:"notion"`
}
