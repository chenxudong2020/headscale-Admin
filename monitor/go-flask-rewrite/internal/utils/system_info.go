package utils

import (
	"math"
	"os"
	"runtime"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type SystemInfo struct {
	CPUUsage           float64 `json:"cpu_usage"`
	MemoryUsagePercent float64 `json:"memory_usage_percent"`
	SentSpeed          float64 `json:"sent_speed"`
	RecvSpeed          float64 `json:"recv_speed"`
}

func GetSystemInfo(netInterface string) (*SystemInfo, error) {
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	netIO, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	var sentSpeed, recvSpeed float64
	for _, iface := range netIO {
		if iface.Name == netInterface {
			sentSpeed = math.Ceil(float64(iface.BytesSent) / 1024)
			recvSpeed = math.Ceil(float64(iface.BytesRecv) / 1024)
			break
		}
	}

	return &SystemInfo{
		CPUUsage:           cpuUsage[0],
		MemoryUsagePercent: memInfo.UsedPercent,
		SentSpeed:          sentSpeed,
		RecvSpeed:          recvSpeed,
	}, nil
}

func GetCurrentDir() (string, error) {
	return os.Getwd()
}