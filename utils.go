package orange

import(
	"fmt"
	"net/http"
)

type HttpError struct {
	Status  int   `json:"status"`
	Message interface{} `json:"message"`
}

// NewHttpError: create http error object
func NewHttpError(status int, message ...interface{}) *HttpError {
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

func concat(s ...string) string {
	size := 0
	for i := 0; i < len(s); i++ {
		size += len(s[i])
	}

	buf := make([]byte, 0, size)

	for i := 0; i < len(s); i++ {
		buf = append(buf, []byte(s[i])...)
	}

	return string(buf)
}
