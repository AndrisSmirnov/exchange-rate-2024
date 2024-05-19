package errors

type ClientError struct {
	Message string `json:"msg"`
}

func NewClientError(message string) *ClientError {
	return &ClientError{Message: message}
}
