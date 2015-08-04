package domain

type User struct {
	Email  string   `json:"email"`
	Name   string   `json:"name"`
	Skills []string `json:"skills"`
}
