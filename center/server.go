package main

import (
	"github.com/nprog/SkyEye/libnet"
	"github.com/nprog/SkyEye/log"
	"github.com/nprog/SkyEye/protocol"
	"github.com/nprog/SkyEye/storage/mongo"
)

type SessionMap map[string]*libnet.Session

type Server struct {
	Options  *ServerOptions
	Sessions *SessionMap
	server   *libnet.Server
	db       *mongo.Mongo
}

func NewServer(c *Config) *Server {
	return &Server{}
}

func (s *Server) handler() error {
	var err error
	return err
}

// // 扫描进程
// func (self *Server) scanDeadSession() {
// 	log.Info("scanDeadSession")
// 	timer := time.NewTicker(self.cfg.ScanDeadSessionTimeout * time.Second)
// 	ttl := time.After(self.cfg.Expire * time.Second)
// 	for {
// 		select {
// 		case <-timer.C:
// 			// log.Info("scanDeadSession timeout")
// 			go func() {
// 			}()
// 		case <-ttl:
// 			break
// 		}
// 	}
// }

func handleSession(s *Server, session *libnet.Session) {
	log.V(1).Info("a new client ", session.Conn().RemoteAddr().String(), " | come in")

	for {
		var msg protocol.Cmd
		if err := session.Receive(&msg); err != nil {
			break
		}

		// err := s.parseProtocol(msg, session)
		// if err != nil {
		// 	log.Error(err.Error())
		// }
	}
}
