package routers

import (
	"gitbub.com/zyl-dev/sgo/api/controllers"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	publicRoute := r.Group("/demo")
	demo := &controllers.Demo{}
	publicRoute.GET("/ping", demo.Ping)

	ginpprof.Wrap(r)
	return r
}
