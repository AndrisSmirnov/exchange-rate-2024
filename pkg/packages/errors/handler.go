package errors

type HandlerError interface {
	HandlingError(error) error
}
