package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
	Password string `json:"password"`
}
