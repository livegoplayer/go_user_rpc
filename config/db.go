package config

import (
	mysqlHelper "github.com/livegoplayer/go_db_helper/mysql"
	"github.com/livegoplayer/go_helper/utils"
	myLogger "github.com/livegoplayer/go_logger/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitDb() {
	databaseConfigMap := viper.GetStringMap("database")

	defaultName := ""
	list := make(map[mysqlHelper.DbName]*mysqlHelper.MysqlConfig, 0)
	for k, m := range databaseConfigMap {
		dbName := mysqlHelper.DbName(k)
		if dbName == "default_db" {
			if dn, ok := m.(utils.H)["name"]; ok && dn != "" {
				defaultName = utils.AsString(dn)
			} else {
				panic("读取默认数据库配置失败")
			}
			continue
		}
		c := &mysqlHelper.MysqlConfig{}
		utils.ToStruct(m.(utils.H), c, "json")
		list[dbName] = c
	}

	mysqlHelper.InitDbList(list, mysqlHelper.DbName(defaultName), myLogger.GetLoggerByLevel(logrus.InfoLevel))
}
