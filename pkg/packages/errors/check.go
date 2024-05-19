package errors

func (err Error) IsCriticalError() bool {
	return err.Type == CriticalError
}
func (err Error) IsDataError() bool {
	return err.Type == DataError
}
func (err Error) IsInternalError() bool {
	return err.Type == InternalError
}
