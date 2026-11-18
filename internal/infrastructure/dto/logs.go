package dto

import "time"

type LinkLog struct {
	Time           time.Time `json:"time"`
	Link           string    `json:"link"`
	ResponseTimeMS int64     `json:"response_time_ms"`
	StatusCode     *int      `json:"status_code"`
	Status         bool      `json:"status"`
}

