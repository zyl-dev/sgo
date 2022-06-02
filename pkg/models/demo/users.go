package demo

import (
	log "github.com/sirupsen/logrus"
	"gitbub.com/zyl-dev/sgo/pkg/common/dbs"
)

type User struct {
	Id        int    `gorm:"id" json:"id"`
	Name      string `gorm:"name" json:"name"`
	CreatedAt string `gorm:"created_at" json:"created_at"` // 创建时间

}

func (t *User) QueryRow() (*User,error) {
	if exit := dbs.CheckDBConnExists(LocalDemoRead); !exit {
		log.WithFields(log.Fields{"func": "GetEmailInfoFromSentLog", "conn": ""}).Error("The Database Not Exists")
	}
	err := dbs.DBMaps[LocalDemoRead].QueryRow("select * from user where id = ?;", 1).Scan(&t.Id, &t.Name, &t.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	return t, err
}