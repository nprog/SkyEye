package main

import (
	"flag"
	"github.com/nprog/SkyEye/common"
	"github.com/nprog/SkyEye/libnet"
)

type ServerOptions struct {
	Port          uint16
	WebPort       uint16
	WebsocketPort uint16
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

	s = NewServer(cfg)

	s.server, err = libnet.Serve("tcp", string(cfg.Server.Port), libnet.Packet(libnet.Uint16BE, libnet.Json()))
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

//
// // 获取内网IP
// func getInternal() {
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		os.Stderr.WriteString("Oops:" + err.Error())
// 		os.Exit(1)
// 	}
//
// 	for _, a := range addrs {
// 		// 过滤内网回环地址
// 		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
// 			if ipnet.IP.To4() != nil {
// 				os.Stdout.WriteString(ipnet.IP.String() + "\n")
// 			}
// 		}
// 	}
// 	os.Exit(0)
// }
