package models

import "errors"

var ErrInvalidPagination = errors.New("limit must be positive and offset must be non-negative")
