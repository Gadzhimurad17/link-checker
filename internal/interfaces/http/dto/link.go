package link

type LinkRequest struct {
	Links []string `json:"links"`
}

type LinkResponse struct {
	Links []LinkResult `json:"links"`
}
type LinkResult struct {
	URL    string `json:"link_url"`
	Status bool
}
