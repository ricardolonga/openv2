package domain

type Event struct {
	Id        string `json:"id,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	Owner     User   `json:"owner,omitempty"`
	Members   []User `json:"members,omitempty"`
}

type Events struct {
	Events []Event `json:"events"`
}
