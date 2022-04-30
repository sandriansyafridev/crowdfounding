package response

import (
	"fmt"
	"strings"
)

func ResponseFail(message string, err interface{}) ApiResponse {

	errorMessage := fmt.Sprint(err)
	if ok := strings.Contains(errorMessage, "\n"); ok {
		split := strings.Split(errorMessage, "\n")
		err = split
	} else {
		err = errorMessage
	}

	return ApiResponse{
		Meta: meta{
			Status:  false,
			Message: message,
		},
		Errors: err,
	}
}
