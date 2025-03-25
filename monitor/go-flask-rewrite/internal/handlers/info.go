package handlers

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

var netInterface = "eth0" // 默认网卡名称

// SetNetInterface 设置网卡名称
func SetNetInterface(interfaceName string) {
	netInterface = interfaceName
}

func GetInfo(c *gin.Context) {
	// 获取 CPU 使用率
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil || len(cpuUsage) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get CPU usage"})
		return
	}

	// 确保 CPU 使用率在 0-100 之间
	cpuUsagePercent := math.Min(math.Max(cpuUsage[0], 0), 100)

	// 获取内存使用率
	memory, err := mem.VirtualMemory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get memory usage"})
		return
	}

	// 确保内存使用率在 0-100 之间
	memoryUsagePercent := math.Min(math.Max(memory.UsedPercent, 0), 100)

	// 获取网卡信息
	netIO, err := net.IOCounters(true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get network IO counters"})
		return
	}

	var sentSpeed, recvSpeed int64
	for _, iface := range netIO {
		if iface.Name == netInterface {
			sentSpeed = int64(math.Ceil(float64(iface.BytesSent) / 1024)) // 转换为 KB
			recvSpeed = int64(math.Ceil(float64(iface.BytesRecv) / 1024)) // 转换为 KB
			break
		}
	}

	// 构建返回数据
	info := gin.H{
		"cpu_usage":            cpuUsagePercent,
		"memory_usage_percent": memoryUsagePercent,
		"sent_speed":           sentSpeed,
		"recv_speed":           recvSpeed,
	}

	c.JSON(http.StatusOK, info)
}
