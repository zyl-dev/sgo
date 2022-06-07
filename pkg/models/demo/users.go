package demo

import (
	"gitbub.com/zyl-dev/sgo/pkg/common/dbs"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Id        int    `gorm:"id" json:"id"`
	Name      string `gorm:"name" json:"name"`
	CreatedAt string `gorm:"created_at" json:"created_at"` // 创建时间

}

func (t *User) QueryRow() {
	var subaccount struct {
		Puid     string `json:"puid"`
		RegionId string `json:"region_id"`
	}
	if exit := dbs.CheckDBConnExists(LocalDemoRead); !exit {
		log.WithFields(log.Fields{"func": "GetEmailInfoFromSentLog", "conn": ""}).Error("The Database Not Exists")
	}
	err := dbs.DBMaps[LocalDemoRead].QueryRow("select puid, region_id from subaccounts where puid = ?;", 643006).Scan(&subaccount.Puid, &subaccount.RegionId)
	if err != nil {
		log.Fatal(err)
	}
}
