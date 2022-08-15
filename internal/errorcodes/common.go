package errorcodes

import "net/http"

type BaseError interface {
	error
	ErrorCode() ErrorType
	ErrorData() map[string]string
}

type ErrorType int

func (e ErrorType) Code() int {
	return int(e)
}

func (e ErrorType) StatusCode() int {
	if statusCode, ok := error2StatusCodeMap[e]; ok {
		return statusCode
	}
	return http.StatusBadRequest
}


const (
	NotFoundError                  ErrorType = 210010001
	UnknownError                   ErrorType = 210010002
	ServerError                    ErrorType = 210010003
	InvalidParamError              ErrorType = 210010004

	DBNotFoundError                   ErrorType = 210020001
	DBOperationError ErrorType = 210020000

	DefaultAppErr   ErrorType = 210030001
	ForbiddenAppErr ErrorType = 210040001

	NOPStatusAppErr                   ErrorType = 210050001

)



var error2StatusCodeMap = map[ErrorType]int{
	InvalidParamError: http.StatusBadRequest,
	DBNotFoundError:   http.StatusBadRequest,
	DBOperationError:  http.StatusBadRequest,
	DefaultAppErr:     http.StatusBadRequest,
	ForbiddenAppErr:   http.StatusBadRequest,
	NOPStatusAppErr:   http.StatusBadRequest,
	ServerError:       http.StatusInternalServerError,
	NotFoundError:     http.StatusNotFound,
}

