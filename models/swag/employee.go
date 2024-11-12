package swag

type CreateUpdateEmployee struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	LogoId      string `json:"logo_id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
