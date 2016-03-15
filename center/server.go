package main

import (
	"github.com/nprog/SkyEye/libnet"
	"github.com/nprog/SkyEye/log"
	"github.com/nprog/SkyEye/protocol"
	"github.com/nprog/SkyEye/storage/mongo"
)

//Server struct
type Server struct {
	Options  *ServerOptions
	Sessions *map[string]*libnet.Session
	server   *libnet.Server
	db       *mongo.Mongo
}

//NewServer server
func NewServer(c *Config) *Server {
	return &Server{}
}

// //scanDeadSession 扫描进程
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

//handleSession 处理连接
func handleSession(s *Server, session *libnet.Session) {
	log.V(1).Info("a new server ", session.Conn().RemoteAddr().String(), " | come in")

	for {
		var msg protocol.Cmd
		if err := session.Receive(&msg); err != nil {
			break
		}

		err := s.parseProtocol(msg, session)
		if err != nil {
			log.Error(err.Error())
		}
	}
}

//parseProtocol 解析协议
func (s *Server) parseProtocol(cmd protocol.Cmd, session *libnet.Session) error {
	var err error

	log.Info(cmd)
	// pp := NewProtoProc(self)
	//
	// cmdName := cmd.GetCmdName()
	//
	// switch cmdName {
	// //PING
	// case protocol.SEND_PING_CMD:
	// 	err = pp.procPing(&cmd, session)
	// 	if err != nil {
	// 		log.Error("error:", err)
	// 		return err
	// 	}
	//
	// default:
	// 	log.Info(cmd.GetCmdName())
	// }
	return err
}
