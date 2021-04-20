package main

import "github.com/livegoplayer/go_db_helper/parse_collect"

var APPROOT = "D:\\files\\workspace\\go\\go_user_rpc"

func main() {
	parseCollect()
}

func parseCollect() {
	parse_collect.Parse(APPROOT)
}
