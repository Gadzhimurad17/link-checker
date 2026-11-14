package dto

import "linkcheker/internal/domain/models"

type LinkRequest struct {
	Links []string `json:"links"`
}

type LinkResponse struct {
	Links []models.Link `json:"links"`
}

type LinkAmountResponse struct {
	Links []models.LinkStat
}
