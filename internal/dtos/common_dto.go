package dtos

const PageSize int64 = 10

type Metadata struct {
	TotalData   int64 `json:"total_data"`
	LastPage    int   `json:"last_page"`
	CurrentPage int   `json:"current_page"`
	PerPage     int64 `json:"perpage"`
}