package response

func ResponseSuccess(message string, data interface{}) ApiResponse {
	return ApiResponse{
		Meta: meta{
			Status:  true,
			Message: message,
		},
		Data: data,
	}
}
