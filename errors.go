package slack

type SlackError struct {
	ErrorType string
}

func NewSlackError(errorType string) SlackError {
	return SlackError{
		ErrorType: errorType,
	}
}

func (e SlackError) Error() string {
	return e.ErrorType
}
