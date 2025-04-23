package uspsgo

type Error struct {
	err ErrorDetails
}

func NewError(err ErrorDetails) Error {
	return Error{err: err}
}

func (e Error) Error() string {
	return e.err.Message
}

func (e Error) Details() ErrorDetails {
	return e.err
}

func (e Error) Is(target error) bool {
	uspsErr, ok := target.(Error)
	if !ok {
		return false
	}
	return e.err.Message == uspsErr.err.Message && e.err.Code == uspsErr.err.Code
}

func (e Error) As(target any) bool {
	uspsErr, ok := target.(*Error)
	if !ok {
		return false
	}
	*uspsErr = e
	return true
}

type ErrorDetails struct {
	APIVersion   string `json:"apiVersion"`
	ErrorMessage `json:"error"`
}

type ErrorMessage struct {
	Code    string               `json:"code"`
	Message string               `json:"message"`
	Errors  []ErrorMessageDetail `json:"errors"`
}

type ErrorMessageDetail struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

var ErrAddressNotFound = Error{
	err: ErrorDetails{
		ErrorMessage: ErrorMessage{
			Code:    "400",
			Message: "Address Not Found",
		},
	},
}
var errAddressNotFoundStr = "Address Not Found"
