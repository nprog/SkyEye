package main

import (
	"github.com/nprog/SkyEye/libnet"
	"github.com/nprog/SkyEye/log"
	"net"
	"os"
)

func main() {
	// flag.Parse()
	getInternal()
	server, err := libnet.Serve("tcp", ":2012", libnet.Packet(libnet.Uint16BE, libnet.Json()))
	if err != nil {
		panic(err)
	}

	for {
		session, err := server.Accept()
		if err != nil {
			break
		}
		log.Info(session)
	}
}

// 获取内网IP
func getInternal() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}

	for _, a := range addrs {
		// 过滤内网回环地址
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	os.Exit(0)
}
