package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/loggers"
)


func NewRouter() *gin.Engine {
	r := gin.New()
	// 	r.Use(cors.Default())
	r.Use(CORSMiddleware())

	r.NoMethod(HandleNotFound)
	r.NoRoute(HandleNotFound)
	r.Use(gin.Recovery())
	r.Use(loggers.GinLogger(logger))

	PublicGroup := r.Group("")
	GroupSysUsers(PublicGroup)

	PrivateGroup := r.Group("")
	PrivateGroup.Use(JWTAuth())
	GroupPreflopRanges(PrivateGroup)
	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			//return
		}

		c.Next()
	}
}