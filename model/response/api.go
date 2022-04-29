package response

type meta struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type api_response struct {
	meta   meta        `json:"meta"`
	data   interface{} `json:"data,omitempty"`
	errors interface{} `json:"errrors,omitempty"`
}
