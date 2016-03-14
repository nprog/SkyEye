package main

import (
	"github.com/nprog/SkyEye/common"
	"github.com/nprog/SkyEye/log"
	// "github.com/shirou/gopsutil/cpu"
	// "github.com/shirou/gopsutil/disk"
	// "github.com/shirou/gopsutil/host"
	// "github.com/shirou/gopsutil/load"
	// "github.com/shirou/gopsutil/mem"
	// "github.com/shirou/gopsutil/net"
	// "github.com/shirou/gopsutil/process"
	"encoding/json"
)

// Test this is test
func Test() {
	rti := common.NewMachineInfo()
	// rti := common.NewRealtimeInfo()
	rti.GetInfo()

	temp, err := json.Marshal(rti)
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info(string(temp))
}
