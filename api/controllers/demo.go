package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gitbub.com/zyl-dev/sgo/pkg/models/demo"
)

type Demo struct {

}


func (t *Demo) Ping(c *gin.Context)   {
	user := &demo.User{}
	res ,err:= user.QueryRow()
	fmt.Println(err)
	c.JSON(http.StatusOK, res)
}
