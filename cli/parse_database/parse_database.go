package main

import (
	"github.com/livegoplayer/go_db_helper/parse_datebase"
	"github.com/livegoplayer/go_helper/config"
	proConfig "github.com/livegoplayer/go_user_rpc/config"
)

var APPROOT = "D:\\files\\workspace\\go\\go_user_rpc"

func main() {
	parseDatabase()
}

func parseDatabase() {
	config.LoadEnv()
	proConfig.InitLog()
	proConfig.InitDb()
	parse_datebase.Parse(APPROOT)
}
