package models

type TV struct {
	Teams  []Team   `json:"teams"`
	TVs    []string `json:"tvs"`
	Time   string   `json:"time"`
	Status string   `json:"status"`
	// Text   string   `json:"test"`
}

type Team struct {
	Name  string `json:"name"`
	Logo  string `json:"logo"`
	Score string `json:"score"`
}
