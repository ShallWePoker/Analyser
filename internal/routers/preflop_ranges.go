package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/configs"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/requests"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/services"
)

func GroupPreflopRanges(g *gin.RouterGroup) {
	group := g.Group(fmt.Sprintf("%s/preflop-ranges", configs.Config.UrlPrefix))

	group.POST("/rfi-ranges", wrapper(generateRFIRanges))
	group.GET("/test-out", wrapper(testOutToken))
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

func testOutToken(c *gin.Context) error {
	return SuccessResp(c, gin.H{"key": "value"})
}
