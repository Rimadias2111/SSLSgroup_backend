package swag

type CreateUpdateCompany struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Number  string `json:"number"`
	SCAC    string `json:"scac"`
	DOT     int    `json:"dot"`
	MC      int    `json:"mc"`
}
