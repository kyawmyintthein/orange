package orange

type HttpError struct{
	Status  int
	Message interface{}
}

// NewHttpError: create http error object
func NewHttpError(code int, message ...interface{}) *HttpError {
	httpError := &HttpError{Code: code, Message: http.StatusText(code)}
	if len(message) > 0 {
		httpError.Message = message[0]
	}
	return httpError
}

// HttpError as string
func (httpError *HttpError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
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