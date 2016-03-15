package main

import (
	"flag"
	"github.com/nprog/SkyEye/common"
	"github.com/nprog/SkyEye/libnet"
	"github.com/nprog/SkyEye/log"
	"os"
)

//ServerOptions Options
type ServerOptions struct {
	ServerAddr string
}

var confFile = flag.String("c", "agent.conf", "conf file.")

func main() {
	var (
		cfg *Config
		c   chan int
	)
	c = make(chan int)

	//VERSION
	version()
	//FLAG PARSE
	flag.Parse()
	//READ CONFIG FILE
	cfg = NewConfig(*confFile)
	// log
	common.SetLogSettings(cfg.Log)
	//TURN ON PPROF
	if cfg.Debug.PprofFile != "" {
		common.Pprof(&cfg.Debug.PprofFile)
	}

	log.Info(cfg.Server)
	log.Info(cfg.Debug)
	log.Info(cfg.Log)

	conn, err := libnet.Connect("tcp", cfg.Server.ServerAddr, libnet.Packet(libnet.Uint16BE, libnet.Json()))
	if err != nil {
		panic(err)
	}

	if err = conn.Send(struct{}{}); err != nil {
		log.Error(err.Error())
	}

	for {
		select {
		case <-c:
			os.Exit(0)
		}
	}
}
