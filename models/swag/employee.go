package swag

type CreateUpdateEmployee struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	LogoId      string `json:"logo_id"`
	Position    string `json:"position"`
	AccessLevel int64  `json:"access_level"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Company     string `json:"company"`
	Birthday    string `json:"birthday"`
	StartDate   string `json:"start_date"`
}
