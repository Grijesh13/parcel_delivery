package dto

type User struct {
	UserName    string `json:"username"`
	Password    string `json:"password"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}
