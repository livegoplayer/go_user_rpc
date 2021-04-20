package main

import "github.com/livegoplayer/go_db_helper/private_model"

var APPROOT = "D:\\files\\workspace\\go\\go_user_rpc"

func main() {
	parsePrivateModel()
}

func parsePrivateModel() {
	private_model.Parse(APPROOT)
}
