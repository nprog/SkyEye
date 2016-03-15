package main

import (
	"github.com/nprog/SkyEye/common"
	"github.com/nprog/SkyEye/libnet"
	"github.com/nprog/SkyEye/log"
	// "github.com/shirou/gopsutil/cpu"
	// "github.com/shirou/gopsutil/disk"
	// "github.com/shirou/gopsutil/host"
	// "github.com/shirou/gopsutil/load"
	// "github.com/shirou/gopsutil/mem"
	// "github.com/shirou/gopsutil/net"
	// "github.com/shirou/gopsutil/process"
	"encoding/json"
	"sync"
	"time"
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

//OnlineCfg cfg
type OnlineCfg struct {
	RealtimeInfo      []string
	RealtimeInfoCycle int64 //更新周期
}

//Collection collection
type Collection struct {
	onlineCfg      *OnlineCfg
	session        *libnet.Session
	changeCfg      bool
	changeCfgMutex sync.Mutex
}

//NewCollection 新建
func NewCollection() *Collection {
	return &Collection{}
}

//sendMachineInfo 发送机器信息
func (c *Collection) sendMachineInfo() {
	log.V(1).Info("sendMachineInfo")
}

//sendRealtimeInfoLoop 循环发送实时数据
func (c *Collection) sendRealtimeInfoLoop() {
	log.Info("sendRealtimeInfoLoop")
	timer := time.NewTicker(5 * time.Second)
	ttl := time.After(10 * time.Second)
bkloop:
	for {
		select {
		case <-timer.C:

			if c.changeCfg {
				c.changeCfg = false
				break bkloop
			}
		case <-ttl:
			break
		}
	}
	log.Info("re config RealtimeInfoLoop.")
	c.sendRealtimeInfoLoop()
}

//sendRealtimeInfoLoop 循环发送实时数据(快速，不入库)
func (c *Collection) sendFastRealtimeInfoLoop() {

}
