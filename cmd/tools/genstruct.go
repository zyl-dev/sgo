package main

import (
	"fmt"
	"gitbub.com/zyl-dev/sgo/pkg/common/dbs"
	"gitbub.com/zyl-dev/sgo/pkg/configs"
	"github.com/gohouse/converter"
	gokitDBS "github.com/zyl-dev/gokit/common/dbs"
)

func init()  {
	configs.InitConfig()
	// init mysql
	mysqlCfg := make(map[string]gokitDBS.MySQLDSN)
	gokitDBS.AddDatabaseConfig(&configs.Config.MySQLDB, mysqlCfg)
	dbs.DBMaps = gokitDBS.DBMaps
}
func main()  {
	err := converter.NewTable2Struct().
		SavePath("./pkg/models/demo/users.go").
		Dsn(configs.Config.MySQLDB.Read.DSN).
		PackageName("demo").
		TagKey("gorm").
		EnableJsonTag(true).
		Table("user").
		Run()
	fmt.Println(err)
}
