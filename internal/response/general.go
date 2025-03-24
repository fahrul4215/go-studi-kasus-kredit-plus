package response

type Success struct {
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Pagination any    `json:"pagination,omitempty"`
	Meta       any    `json:"meta,omitempty"`
}
