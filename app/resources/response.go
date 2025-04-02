package resources

type Response struct {
	Message string `json:"messsage"`
	Status  bool   `json:"status"`
	Data    any    `json:"data,omitempty"`
}
