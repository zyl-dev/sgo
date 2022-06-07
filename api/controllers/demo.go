package controllers

import (
	"net/http"

	"gitbub.com/zyl-dev/sgo/pkg/models/demo"
	"github.com/gin-gonic/gin"
)

type Demo struct {
}

func (t *Demo) Ping(c *gin.Context) {
	user := &demo.User{}
	user.QueryRow()

	c.JSON(http.StatusOK, "success")
}
