package main

import (
	"flag"
	"github.com/nprog/SkyEye/common"
	"github.com/nprog/SkyEye/libnet"
	"github.com/nprog/SkyEye/log"
)

type ServerOptions struct {
	Port          string
	WebPort       string
	WebsocketPort string
	Daemon        bool
}

var confFile = flag.String("c", "center.conf", "conf file.")

func main() {
	var (
		cfg *Config
		c   chan int
		err error
		s   *Server
	)
	//VERSION
	version()
	//FLAG PARSE
	flag.Parse()
	//READ CONFIG FILE
	cfg = NewConfig(*confFile)
	//Daemon
	common.Daemon(cfg.Server.Daemon)
	// log
	common.SetLogSettings(cfg.Log)
	//TURN ON PPROF
	if cfg.Debug.PprofFile != "" {
		common.Pprof(&cfg.Debug.PprofFile)
	}

	log.Info(cfg.Db)
	log.Info(cfg.Debug)
	log.Info(cfg.Log)
	log.Info(cfg.Server)

	s = NewServer(cfg)

	s.server, err = libnet.Serve("tcp", cfg.Server.Port, libnet.Packet(libnet.Uint16BE, libnet.Json()))
	if err != nil {
		panic(err)
	}

	for {
		session, err := s.server.Accept()
		if err != nil {
			break
		}

		go handleSession(s, session)
	}

	select {
	case <-c:
		break
	}
}
