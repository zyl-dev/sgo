package controllers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"gitbub.com/zyl-dev/sgo/pkg/common/dbs"

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

func (t *Demo) RedisPing(c *gin.Context) {

	//err := dbs.RedisDB.Get("test_key").Scan(&val)
	val := dbs.RedisDB.Get("testd_key").Val()
	if val == "" {
		log.Error("key testd_key not exist!")
	}
	fmt.Println(val)
}
