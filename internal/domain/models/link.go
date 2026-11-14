package models

type LinkStat struct {
	URL    string `json:"url"`
	Amount uint64 `json:"amount"`
}
type Link struct {
	URL    string `json:"link"`
	Status bool   `json:"status"`
}

type LinksStats struct {
	Stats []LinkStat `json:"stats"`
}
