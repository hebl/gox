package xweb

//ErrorType Error Type
type ErrorType uint64

type (
	//Error Error info
	Error struct {
		Err  error
		Type ErrorType
		Meta interface{}
	}

	errorMsgs []*Error
)

var _ error = &Error{}

// SetType SetType
func (msg *Error) SetType(flags ErrorType) *Error {
	msg.Type = flags
	return msg
}

// SetMeta SetMeta
func (msg *Error) SetMeta(data interface{}) *Error {
	msg.Meta = data
	return msg
}

// Implements the error interface
func (msg *Error) Error() string {
	return msg.Err.Error()
}
