package model

import (
	"framework/config"

	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
)

//var db *gdb.DB
//var tablePrefix string

func init() {
	//tablePrefix = GetPrefix()
	db := g.DB()
	db.SetMaxIdleConnCount(10)
	db.SetMaxOpenConnCount(100)
	db.SetDebug(config.GetCfg().GetBool("run_debug", true))
}

type modelType interface {
	TableName() string
}

func GetDB() gdb.DB {
	//return db
	return g.DB()
}

func GetPrefix() string {
	return config.GetCfg().GetString("database.prefix", "")
}

func Model(mod modelType) *gdb.Model {
	return GetDB().Table(mod.TableName()).Safe()
}
