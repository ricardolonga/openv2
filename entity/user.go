package entity

type User struct {
	Email  string   `json:"email,omitempty"`
	Name   string   `json:"name,omitempty"`
	Skills []string `json:"skills,omitempty"`
}
