package routers

import (
	"github.com/gin-gonic/gin"
	"gitbub.com/zyl-dev/sgo/api/controllers"
)
import "github.com/DeanThompson/ginpprof"
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