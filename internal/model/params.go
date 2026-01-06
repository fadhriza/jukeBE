package model

type PaginationQuery struct {
	Page      int
	Limit     int
	SortBy    string
	SortOrder string
	Search    string
	// Filters
	Position string
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int         `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"total_pages"`
}
