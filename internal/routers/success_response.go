package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/responses"
	"net/http"
)

func SuccessResp(c *gin.Context, data interface{}) error {
	res := responses.APISuccessResp{
		Code: responses.SUCCESS_CODE,
		Msg:  "success",
		Data: data,
	}
	c.JSON(http.StatusOK, res)
	return nil
}
