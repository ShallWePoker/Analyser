package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/requests"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/services"
)

func GroupPreflopRanges(r *gin.Engine) {
	// TODO api/v1 read from configs
	group := r.Group(fmt.Sprintf("%s/preflop-ranges", "api/v1"))

	group.POST("/rfi-ranges", wrapper(generateRFIRanges))
}

func generateRFIRanges(c *gin.Context) error {
	var txtFilePathReq requests.GGTXTFilePathReq
	err := c.ShouldBindJSON(&txtFilePathReq)
	if err != nil {
		return err
	}
	resp, err := services.GenerateRFIRanges(txtFilePathReq.TXTFilePath)
	if err != nil {
		return err
	}
	return SuccessResp(c, resp)
}
