package main

import (
	"gitbub.com/zyl-dev/sgo/api/routers"
	"gitbub.com/zyl-dev/sgo/pkg/common/dbs"
	"gitbub.com/zyl-dev/sgo/pkg/configs"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	log "github.com/sirupsen/logrus"
	gokitDBS "github.com/zyl-dev/gokit/common/dbs"
	gokitLogger "github.com/zyl-dev/gokit/common/logger"
)

func init()  {
	configs.InitConfig()
	// init mysql
	mysqlCfg := make(map[string]gokitDBS.MySQLDSN)
	gokitDBS.AddDatabaseConfig(&configs.Config.MySQLDB, mysqlCfg)
	gokitDBS.InitMySQLDB(mysqlCfg, configs.Config.SSHConfigs)
	dbs.DBMaps = gokitDBS.DBMaps
	// init redis
	_ = gokitDBS.InitRedisCacheDB(&configs.Config.Redis)
	dbs.RedisDB = gokitDBS.RedisDB
	// init log
	gokitLogger.InitLog(configs.Config.Log.Level, configs.Config.Log.Path, "gin")
}
func main() {
	if configs.Config.Pyroscope.IsEnabled && configs.Config.Pyroscope.Endpoint != "" {
		_, _ = profiler.Start(profiler.Config{
			// TODO: add SampleRate, https://github.com/pyroscope-io/pyroscope/blob/main/pkg/agent/profiler/profiler.go
			ApplicationName: "pushGateway-api",
			ServerAddress: configs.Config.Pyroscope.Endpoint,
			ProfileTypes: []profiler.ProfileType{
				profiler.ProfileCPU,
				profiler.ProfileAllocObjects,
				profiler.ProfileAllocSpace,
				profiler.ProfileInuseObjects,
				profiler.ProfileInuseSpace,
			},
			DisableGCRuns: true,
		})
	}
	r := routers.InitRouter()
	err := r.Run()
	if err != nil {
		log.Error(err)
	}
	defer func() {
		log.Debugf("close producer daemon")
	}()
}