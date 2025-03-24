package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type InfoResponse struct {
	CPUUsage           float64 `json:"cpu_usage"`
	MemoryUsagePercent float64 `json:"memory_usage_percent"`
	SentSpeed          uint64  `json:"sent_speed"`
	RecvSpeed          uint64  `json:"recv_speed"`
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	cpuUsage, _ := cpu.Percent(0, false)
	memInfo, _ := mem.VirtualMemory()
	netIO, _ := net.IOCounters(false)

	var sentSpeed, recvSpeed uint64
	if len(netIO) > 0 {
		sentSpeed = netIO[0].BytesSent
		recvSpeed = netIO[0].BytesRecv
	}

	infoResponse := InfoResponse{
		CPUUsage:           cpuUsage[0],
		MemoryUsagePercent: memInfo.UsedPercent,
		SentSpeed:          sentSpeed,
		RecvSpeed:          recvSpeed,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(infoResponse)
}