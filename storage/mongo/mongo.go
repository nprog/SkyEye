package mongo

import (
	"github.com/nprog/IntelligentEngine/glog"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"strconv"
	"sync"
	"time"
)

type MongoOptions struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DbName   string
	Uri      string
}

type Mongo struct {
	Opts    *MongoOptions
	Session *mgo.Session

	rwMutex      sync.Mutex
	ReConnecting bool
}

func NewMongo(opts *MongoOptions) *Mongo {
	if opts.Uri == "" {
		var uri string
		if opts.User == "" && opts.Password == "" {
			uri = opts.Host + ":" + strconv.Itoa(int(opts.Port))
		} else {
			uri = opts.User + ":" + opts.Password + "@" + opts.Host + ":" +
				strconv.Itoa(int(opts.Port))
		}

		opts.Uri = uri
	}
	if opts.DbName == "" {
		opts.DbName = DEFAULT_DB_NAME
	}

	glog.Info("connect to mongo : ", opts.Uri)
	maxWait := time.Duration(5 * time.Second)
	session, err := mgo.DialWithTimeout(opts.Uri, maxWait)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	return &Mongo{
		Opts:         opts,
		Session:      session,
		ReConnecting: false,
	}
}

//Close
func (m *Mongo) Close() {
	m.Session.Close()
}

//ReConnect
func (m *Mongo) ReConnect() {
	glog.Info("ReConnecting")
	maxWait := time.Duration(5 * time.Second)
	session, err := mgo.DialWithTimeout(m.Opts.Uri, maxWait)

	if err != nil {
		m.ReConnecting = false
		glog.Error(err.Error())
		return
	}

	session.SetMode(mgo.Monotonic, true)
	m.Session = session
	m.ReConnecting = false
}

//统一错误处理
func (m *Mongo) Error(err error) error {
	if err.Error() == "EOF" {
		if m.ReConnecting == false {
			m.ReConnecting = true
			m.ReConnect()
		}
	}
	return err
}
