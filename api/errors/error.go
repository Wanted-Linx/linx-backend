package errors

var (
	BadRequestErr   error
	ForbiddenErr    error
	UnauthorizedErr error
	NotFoundErr     error
	InternalErr     error
)

type RestErr struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	ErrorType string `json:"error_type"`
}

func NewRestError(code int, message, errorType string) *RestErr {
	return &RestErr{
		Code:      code,
		Message:   message,
		ErrorType: errorType,
	}
}

func (e *RestErr) Error() string {
	return e.ErrorType + ":" + e.Message
}
