package main

import (
	"flag"
	"fmt"
	"go-flask-rewrite/internal/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	netInterface string // 全局变量，用于存储网卡名称
)

func main() {
	// 定义命令行参数
	flag.StringVar(&netInterface, "interface", "eth0", "Network interface to monitor (default: eth0)")
	flag.Parse()

	// 处理 help 参数
	if len(os.Args) > 1 && os.Args[1] == "help" {
		fmt.Println("Available parameters:")
		fmt.Println("  -interface: Network interface to monitor (default: eth0)")
		os.Exit(0)
	}

	// 设置 Gin 的日志模式为 release 模式
	gin.SetMode(gin.ReleaseMode)

	// 将网卡名称传递给 handlers
	handlers.SetNetInterface(netInterface)

	// 初始化 Gin 路由
	r := gin.Default()

	// 定义路由
	r.GET("/info", handlers.GetInfo)
	r.GET("/reload_acl", handlers.ReloadACL)
	r.GET("/data_record", handlers.DataRecord)
	r.GET("/data_usage", handlers.DataUsage)

	// 启动服务
	r.Run(":5000") // 监听 5000 端口
}
