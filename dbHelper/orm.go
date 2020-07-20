package dbHelper

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//这里存放的是一些封装
type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int32
	Dbname   string
}

var mysqlConfig *MysqlConfig
var _db *gorm.DB

func InitDbHelper(mysqlCfg *MysqlConfig, logMode bool, MaxOpenCon int, MaxIdleCon int) {
	//初始化全局sql连接

	mysqlConfig = mysqlCfg
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Dbname)
	//连接MYSQL
	db, err := gorm.Open("mysql", mysqlDsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	_db = db

	//打开调试模式
	db.LogMode(logMode)

	//设置数据库连接池参数
	_db.DB().SetMaxOpenConns(MaxOpenCon) //设置数据库连接池最大连接数
	_db.DB().SetMaxIdleConns(MaxIdleCon) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return _db
}
