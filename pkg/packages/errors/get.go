package errors

func (err Error) GetMessage() string {
	if err.Message != "" {
		return err.Message
	}
	if err.ExtendedError.Error() != "" {
		return err.ExtendedError.Error()
	}
	return "no message"
}

func (err Error) GetCode() int {
	if err.Code != 0 {
		return err.Code
	}
	return err.Type.ToCode()
}
