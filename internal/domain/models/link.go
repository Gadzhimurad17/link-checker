package models

type LinkStat struct {
	URL    string `json:"url"`
	Amount uint64 `json:"amount"`
}

type LinksStats struct {
	Stats []LinkStat `json:"stats"`
}
