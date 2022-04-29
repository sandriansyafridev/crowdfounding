package response

import (
	"fmt"
	"strings"
)

func ResponseFail(message string, err interface{}) api_response {

	errorMessage := fmt.Sprint(err)
	ok := strings.Contains(errorMessage, "\n")
	if ok {
		split := strings.Split(errorMessage, "\n")
		err = split
	} else {
		err = errorMessage
	}

	return api_response{
		meta: meta{
			Status:  false,
			Message: message,
		},
		errors: err,
	}
}
