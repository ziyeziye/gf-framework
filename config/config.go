package config

import (
	"framework/library/utli"
	"strings"

	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/gcfg"
)

var (
	conf *gcfg.Config
)

func init() {
	conf = g.Config()
	confFile := "config.default.toml"
	//conf.SetPath(path)
	conf.SetFileName(confFile)
}

func GetRealPath(keyName string, def ...interface{}) string {
	return GetRootPath() + "/" + strings.TrimLeft(GetCfg().GetString(keyName, def), "/")
}

func GetRootPath() string {
	filePath, _ := utli.CurrentFilePath()
	return utli.Dirname(utli.Dirname(filePath))
}

func GetCfg() *gcfg.Config {
	return conf
}
