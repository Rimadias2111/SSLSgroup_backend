package swag

type CreateUpdateDriver struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	TruckNumber string `json:"truck_number"`
	PhoneNumber string `json:"phone_number"`
	Mail        string `json:"mail"`
	Birthday    string `json:"birthday"`
	CompanyId   string `json:"company_id"`
	StartDate   string `json:"start_date"`
	Type        string `json:"type"`
	Position    string `json:"position"`
}
