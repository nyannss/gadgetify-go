package base

type MetaPagination struct {
	TotalData   uint32 `json:"total_data"`
	TotalPage   uint16 `json:"total_page"`
	CurrentPage uint16 `json:"current_page"`
	PerPage     uint   `json:"per_page"`
}

type BasePagination struct {
	Meta  MetaPagination `json:"meta"`
	Items interface{}    `json:"items"`
}
