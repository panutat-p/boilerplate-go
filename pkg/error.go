package pkg

type HttpError struct {
	Err        error
	StatusCode int
	Message    string
}

func (e *HttpError) Error() string {
	return e.Err.Error()
}
