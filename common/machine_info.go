package common

import (
	"github.com/nprog/SkyEye/log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"runtime"
)

// MachineInfo 系统基本信息
type MachineInfo struct {
	CPU           *[]cpu.CPUInfoStat        `json:"cpu"`
	Mem           *mem.VirtualMemoryStat    `json:"mem"`
	Swap          *mem.SwapMemoryStat       `json:"swap"`
	Host          *host.HostInfoStat        `json:"host"`
	DiskUsage     *disk.DiskUsageStat       `json:"disk_usage"`
	DiskPartition *[]disk.DiskPartitionStat `json:"disk_parititon"`
	Net           *[]net.NetInterfaceStat   `json:"net"`
}

// NewMachineInfo create
func NewMachineInfo() *MachineInfo {
	return &MachineInfo{}
}

// GetInfo 执行获取机器基本信息的函数
func (m *MachineInfo) GetInfo() {
	m.GetCPUInfo()
	m.GetMemInfo()
	m.GetSwapInfo()
	m.GetHostInfo()
	m.GetDiskUsageInfo()
	m.GetDiskPartitionInfo()
	m.GetNetInfo()
}

// GetCPUInfo 获取CPU基本信息
func (m *MachineInfo) GetCPUInfo() {
	cpuInfo, err := cpu.CPUInfo()
	if err != nil {
		log.Error(err.Error())
	}
	m.CPU = &cpuInfo
}

// GetMemInfo 获取内存基本信息
func (m *MachineInfo) GetMemInfo() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Error(err.Error())
	}
	m.Mem = memInfo
}

// GetSwapInfo 获取交换分区信息
func (m *MachineInfo) GetSwapInfo() {
	swapInfo, err := mem.SwapMemory()
	if err != nil {
		log.Error(err.Error())
	}
	m.Swap = swapInfo
}

// GetHostInfo 获取主机基本信息
func (m *MachineInfo) GetHostInfo() {
	hostInfo, err := host.HostInfo()
	if err != nil {
		log.Error(err.Error())
	}
	m.Host = hostInfo
}

// GetDiskUsageInfo 获取硬盘基本信息
func (m *MachineInfo) GetDiskUsageInfo() {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	diskUsageInfo, err := disk.DiskUsage(path)
	if err != nil {
		log.Error(err.Error())
	}
	m.DiskUsage = diskUsageInfo
}

// GetDiskPartitionInfo 获取硬盘分区基本信息
func (m *MachineInfo) GetDiskPartitionInfo() {
	diskPartitionInfo, err := disk.DiskPartitions(true)
	if err != nil {
		log.Error(err.Error())
	}
	m.DiskPartition = &diskPartitionInfo
}

// GetNetInfo 获取网络基本信息
func (m *MachineInfo) GetNetInfo() {
	netInfo, err := net.NetInterfaces()
	if err != nil {
		log.Error(err.Error())
	}
	m.Net = &netInfo
}
