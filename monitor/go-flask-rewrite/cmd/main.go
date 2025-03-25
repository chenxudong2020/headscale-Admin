package main

import (
	"flag"
	"go-flask-rewrite/internal/handlers"

	"github.com/gin-gonic/gin"
)

var netInterface string // 全局变量，用于存储网卡名称

func main() {
	// 定义命令行参数 interface，默认值为 eth0
	flag.StringVar(&netInterface, "interface", "eth0", "Network interface to monitor (default: eth0)")
	flag.Parse()

	// 将网卡名称传递给 handlers
	handlers.SetNetInterface(netInterface)

	r := gin.Default()

	r.GET("/info", handlers.GetInfo)
	r.GET("/reload_acl", handlers.ReloadACL)
	r.GET("/data_record", handlers.DataRecord)
	r.GET("/data_usage", handlers.DataUsage)

	r.Run(":5000") // 监听 5000 端口
}
