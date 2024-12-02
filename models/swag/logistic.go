package swag

type CreateUpdateLogistic struct {
	DriverId string `json:"driver_id"`
	CargoId  string `json:"cargo_id"`
	Status   string `json:"status"`
	StTime   string `json:"st_time"`
	Location string `json:"location"`
	Notion   string `json:"notion"`
	Post     bool   `json:"post"`
}

type UpdateLogisticWithCargo struct {
	LogisticId   string  `json:"logistic_id"`
	Status       string  `json:"status"`
	CargoId      string  `json:"cargo_id"`
	Notion       string  `json:"notion"`
	StTime       string  `json:"st_time"`
	Location     string  `json:"location"`
	Post         bool    `json:"post"`
	LoadId       string  `json:"load_id"`
	Provider     string  `json:"provider"`
	LoadedMiles  int64   `json:"loaded_miles"`
	FreeMiles    int64   `json:"free_miles"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	Cost         int64   `json:"cost"`
	Rate         float64 `json:"rate"`
	PickUpTime   string  `json:"pick_up_time"`
	DeliveryTime string  `json:"delivery_time"`
	EmployeeId   string  `json:"employee_id"`
	Create       bool    `json:"create"`
}

type TerminateLogistic struct {
	LogisticId string `json:"logistic_id"`
	Success    bool   `json:"success"`
}

type CancelLogistic struct {
	LogisticId string `json:"logistic_id"`
	Cancel     bool   `json:"cancel"`
	WhoseFault string `json:"whose_fault"`
	Status     string `json:"status"`
	Section    string `json:"section"`
	Reason     string `json:"reason"`
	CompanyId  string `json:"company_id"`
}
