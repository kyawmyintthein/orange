package orange

import(
	"fmt"
	"net/http"
)

// Error
var (
	notFoundError       = newHttpError(http.StatusNotFound)
	internalServerError = newHttpError(http.StatusInternalServerError)
)

type HttpError struct {
	Status  int   `json:"status"`
	Message interface{} `json:"message"`
}

// newHttpError: create http error object
func newHttpError(status int, message ...interface{}) *HttpError {
	httpError := &HttpError{Status: status, Message: http.StatusText(status)}
	if len(message) > 0 {
		httpError.Message = message[0]
	}
	return httpError
}

// HttpError as string
func (httpError *HttpError) Error() string {
	return fmt.Sprintf("status=%d, message=%v", httpError.Status, httpError.Message)
}
