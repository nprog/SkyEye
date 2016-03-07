package common

import (
	"github.com/nprog/SkyEye/log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// RealtimeInfo 实时系统状态
type RealtimeInfo struct {
	CPU     *[]cpu.CPUTimesStat                 `json:"cpu"`
	Mem     *mem.VirtualMemoryStat              `json:"mem"`
	Swap    *mem.SwapMemoryStat                 `json:"swap"`
	Net     *[]net.NetIOCountersStat            `json:"net"`
	DiskIO  *map[string]disk.DiskIOCountersStat `json:"disk_io"`
	Load    *load.LoadAvgStat                   `json:"load"`
	Process *[]process.Process                  `json:"process"`
}

// NewRealtimeInfo 新建实时系统状态
func NewRealtimeInfo() *RealtimeInfo {
	return &RealtimeInfo{}
}

// GetInfo 执行获取机器当前状态的函数
func (r *RealtimeInfo) GetInfo() {
	r.GetCPUInfo()
	r.GetMemInfo()
	r.GetSwapInfo()
	r.GetNetInfo()
	r.GetDiskIOInfo()
	r.GetLoadInfo()
	// r.GetProcessInfo()
}

// GetCPUInfo 获取CPU实时信息
func (r *RealtimeInfo) GetCPUInfo() {
	cpuInfo, err := cpu.CPUTimes(false)
	if err != nil {
		log.Error(err.Error())
	}
	r.CPU = &cpuInfo
}

// GetMemInfo 获取内存实时信息
func (r *RealtimeInfo) GetMemInfo() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Error(err.Error())
	}
	r.Mem = memInfo
}

// GetSwapInfo 获取交换空前实时信息
func (r *RealtimeInfo) GetSwapInfo() {
	swapInfo, err := mem.SwapMemory()
	if err != nil {
		log.Error(err.Error())
	}
	r.Swap = swapInfo
}

// GetNetInfo 获取网络实时信息
func (r *RealtimeInfo) GetNetInfo() {
	netInfo, err := net.NetIOCounters(false)
	if err != nil {
		log.Error(err.Error())
	}
	r.Net = &netInfo
}

// GetDiskIOInfo 获取磁盘实时信息
func (r *RealtimeInfo) GetDiskIOInfo() {
	diskIOInfo, err := disk.DiskIOCounters()
	if err != nil {
		log.Error(err.Error())
	}
	r.DiskIO = &diskIOInfo
}

// GetLoadInfo 获取Load文件实时信息(Linux)
func (r *RealtimeInfo) GetLoadInfo() {
	loadInfo, err := load.LoadAvg()
	if err != nil {
		log.Error(err.Error())
	}
	r.Load = loadInfo
}

// // GetProcessInfo 获取进程实时信息
// func (r *RealtimeInfo) GetProcessInfo() {
// 	processInfo, err := process.CPUInfo()
// 	if err != nil {
// 		log.Error(err.Error())
// 	}
// 	r.CPU = &processInfo
// }
