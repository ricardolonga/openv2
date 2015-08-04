package domain

type Event struct {
	Id        string `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Owner     User   `json:"owner"`
	Members   []User `json:"members"`
}

type Events struct {
	Events []Event `json:"events"`
}
