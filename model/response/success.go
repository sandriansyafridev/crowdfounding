package response

func ResponseSuccess(message string, data interface{}) api_response {
	return api_response{
		meta: meta{
			Status:  true,
			Message: message,
		},
		data: data,
	}
}
