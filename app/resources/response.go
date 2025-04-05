package resources

type Response struct {
	Message   string `json:"messsage"`
	Status    bool   `json:"status"`
	Data      any    `json:"data,omitempty"`
	Code      int    `json:"code,omitempty"`
	Duplicate bool   `json:"duplicate,omitempty"`
	Total     int    `json:"total,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}
