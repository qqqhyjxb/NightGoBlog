package main

import (
	"NightGoBlog/model"
	"NightGoBlog/routes"
)

func main() {
	// 引用数据库
	model.InitDb()

	routes.InitRouter()
}
