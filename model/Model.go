package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	AddDatetime    string `gorm:"column:add_datetime;-"`
	UpdateDatetime string `gorm:"column:update_datetime;-"`
}

var (
	prefix string
)

func init() {
	prefix = "fs_"
}
