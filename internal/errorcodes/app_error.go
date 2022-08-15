package errorcodes

type AppError struct {
	ErrorType ErrorType
	ErrorMsg  string
	Data      map[string]string
}

func (err *AppError) Error() string {
	if err.ErrorMsg != "" {
		return err.ErrorMsg
	}
	switch err.ErrorType {
	case DefaultAppErr:
		return "Bad request!"
	case ForbiddenAppErr:
		return "Permission denied!"
	case NOPStatusAppErr:
		return "Current status cannot be operated!"
	case NotFoundError:
		return "Not found!"
	case ServerError:
		return "Internal server error!"
	case UnknownError:
		return "Unknown error"
	default:
		return "Bad request!"
	}
}


func (err *AppError) ErrorCode() ErrorType {
	return err.ErrorType
}

func (err *AppError) ErrorData() map[string]string {
	return err.Data
}

func NewDefaultAppErr(msg string) *AppError {
	return &AppError{
		ErrorMsg:  msg,
		ErrorType: DefaultAppErr,
	}
}

func NewNotFoundAppErr() *AppError {
	return &AppError{
		ErrorType: NotFoundError,
	}
}

