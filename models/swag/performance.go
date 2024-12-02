package swag

type CreateUpdatePerformance struct {
	Reason     string `json:"reason"`
	WhoseFault string `json:"whose_fault"`
	Status     string `json:"status"`
	Section    string `json:"section"`
	EmployeeId string `json:"employee_id"`
	CompanyId  string `json:"company_id"`
	LoadId     string `json:"load_id"`
}
