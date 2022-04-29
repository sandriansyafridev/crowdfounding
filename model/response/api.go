package response

type meta struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ApiResponse struct {
	Meta   meta        `json:"meta"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errrors,omitempty"`
}
