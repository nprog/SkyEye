package main

import (
	"github.com/nprog/SkyEye/libnet"
	"github.com/nprog/SkyEye/log"
)

func main() {
	c, err := libnet.Connect("tcp", "127.0.0.1:2012", libnet.Packet(libnet.Uint16BE, libnet.Json()))
	if err != nil {
		panic(err)
	}

	if err = c.Send(struct{}{}); err != nil {
		log.Error(err.Error())
	}
	// Test()
}
