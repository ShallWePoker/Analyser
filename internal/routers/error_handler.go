package routers

import (
	"github.com/gin-gonic/gin"
	. "github.com/shallwepoker/ggpoker-hands-converter/internal/errorcodes"
)

type HandlerFunc func(c *gin.Context) error

type APIException struct {
	StatusCode int               `json:"-"`
	Code       int               `json:"code"`
	Msg        string            `json:"msg"`
	Data       map[string]string `json:"data"`
}

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := handler(c)
		if err == nil {
			return
		}
		apiException := handleErr(err)
		c.JSON(apiException.StatusCode, apiException)
	}
}


func handleErr(err error) *APIException {
	var apiException *APIException
	switch err.(type) {
	case BaseError:
		apiException = convertToApiException(err.(BaseError))
	case error:
		apiException = convertToApiException(NewDefaultAppErr(err.Error()))
	}
	return apiException
}

func convertToApiException(err BaseError) *APIException {
	return &APIException{
		StatusCode: err.ErrorCode().StatusCode(),
		Code:       err.ErrorCode().Code(),
		Msg:        err.Error(),
		Data:       err.ErrorData(),
	}
}


func HandleNotFound(c *gin.Context) {
	err := NewNotFoundAppErr()
	exception := handleErr(err)
	c.JSON(exception.StatusCode, exception)
	return
}
