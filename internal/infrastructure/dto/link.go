package dto

type LinkRequest struct {
	Links []string `json:"links"`
}

type LinkResponse struct {
	Links []LinkResult `json:"links"`
}

type LinkResult struct {
	URL    string `json:"link"`
	Status bool   `json:"status"`
}

type LinkAmountResult struct {
	URL    string `json:"link"`
	Amount int    `json:"amount"`
}
type LinkAmountResponse struct {
	Links []LinkAmountResult 
}
