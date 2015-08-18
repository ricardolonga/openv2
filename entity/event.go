package entity

type Event struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	Owner     User   `json:"owner,omitempty"`
	Members   []User `json:"members,omitempty"`
}

