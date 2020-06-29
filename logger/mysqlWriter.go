package logger

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"go_user_rpc/helper"
)

//手动实现一个mysql_writer,作为输出流对象传递到log
type MysqlWriter struct {
	DbConnection *sql.DB
}

//设计mongo日志表存储格式

var mw *MysqlWriter = &MysqlWriter{}

func init() {
	var err error
	mw.DbConnection, err = sql.Open("mysql", "myuser:myuser@tcp(139.224.132.234:3306)/file_store")
	if err != nil {
		fmt.Printf("err:" + err.Error())
	}
	//设置数据库最大连接数
	mw.DbConnection.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	mw.DbConnection.SetMaxIdleConns(10)
	//验证连接
	if err := mw.DbConnection.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

func (mw *MysqlWriter) Write(p []byte) (n int, err error) {
	n = 0
	err = nil

	//解析出level_no
	levelNo := helper.GetJsonVal(p, "levelNo")

	//开启事务
	tx, err := mw.DbConnection.Begin()
	if err != nil {
		fmt.Println("获取事务失败")
		return
	}

	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO `fs_log` (`message`, `level`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return
	}

	//将参数传递到sql语句中并且执行
	_, err = stmt.Exec(p, levelNo)
	if err != nil {
		fmt.Println("Exec fail")
		return
	}

	//将事务提交
	err = tx.Commit()
	//获得上一个插入自增的id
	if err != nil {
		fmt.Println("commit fail")
		return
	}

	return len(p), nil
}
