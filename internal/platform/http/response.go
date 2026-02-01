package http

type Meta struct {
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
}

type PagedResponse struct {
	Data any  `json:"data"`
	Meta Meta `json:"meta"`
}
