package configs

import (
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
	gokitDBS "github.com/zyl-dev/gokit/common/dbs"
	gokitSSH "github.com/zyl-dev/gokit/common/ssh"
)
var Config = struct {
	AppName string
	Log                    Log
	Pyroscope              Pyroscope
	MySQLDB                gokitDBS.MySQLDB
	Redis                  gokitDBS.RedisConfig
	SSHConfigs             map[string]gokitSSH.SSHConfig
}{}
// Log 代表日志配置
type Log struct {
	Level      string
	ShowSource bool
	Path       string
}

// Pyroscope
type Pyroscope struct {
	IsEnabled bool
	Endpoint  string // Server 地址
}

func InitConfig()  {
	if err := configor.Load(&Config, "configs/config.yaml"); err != nil {
		panic(errors.Wrapf(err, "load config file failed"))
	}
}