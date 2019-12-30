package models

// User type
type User struct {
	ID       int    `json:"id,omiempty"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omiempty"`
	Phone    string `json:"phone,omiempty"`
}
