package main

import (
	"github.com/Terry-Mao/goconf"
	"github.com/nprog/SkyEye/common"
	"github.com/nprog/SkyEye/log"
	"github.com/nprog/SkyEye/storage/mongo"
	"strconv"
)

//Config 读取配置
type Config struct {
	Server *ServerOptions
	Log    *common.LogOptions
	Debug  *common.DebugOptions
	Db     *mongo.MongoOptions
}

//ConfigTemp 文件格式
type ConfigTemp struct {
	//server
	ServerPort          uint16 `goconf:"server:port"`
	ServerWebPort       uint16 `goconf:"server:web_port"`
	ServerWebsocketPort uint16 `goconf:"server:websocket_port"`
	ServerDaemon        string `goconf:"server:daemon"`

	//log
	LogLogtofile string `goconf:"log:logtofile"`
	LogDir       string `goconf:"log:dir"`
	LogLogtostd  string `goconf:"log:logtostd"`
	LogVerbosity string `goconf:"log:verbosity"`

	//db
	DbHost     string `goconf:"db:host"`
	DbPort     uint16 `goconf:"db:port"`
	DbUser     string `goconf:"db:user"`
	DbPassword string `goconf:"db:password"`
	DbName     string `goconf:"db:name"`
	DbUri      string `goconf:"db:uri"`

	//debug
	DebugPprofFile string `goconf:"debug:pprof_file"`
}

//NewConfig Config
func NewConfig(confFile string) *Config {
	temp := &ConfigTemp{}

	conf := goconf.New()
	if err := conf.Parse(confFile); err != nil {
		log.Error(err.Error())
		return temp.Parse()
	}

	// temp := &ConfigTemp{}
	if err := conf.Unmarshal(temp); err != nil {
		log.Error(err.Error())
		return temp.Parse()
	}

	return temp.Parse()
}

//Parse 解析配置
func (c *ConfigTemp) Parse() *Config {
	var (
		conf          Config
		serverOptions ServerOptions
		logOptions    common.LogOptions
		mongoOptions  mongo.MongoOptions
		debugOptions  common.DebugOptions
	)

	//Server
	if c.ServerPort > MIN_PORT && c.ServerPort < MAX_PORT {
		serverOptions.Port = ":" + strconv.Itoa(int(c.ServerPort))
	} else {
		serverOptions.Port = CONFIG_DEFAULT_SERVER_PORT
	}
	if c.ServerWebPort > MIN_PORT && c.ServerWebPort < MAX_PORT {
		serverOptions.WebPort = ":" + strconv.Itoa(int(c.ServerWebPort))
	} else {
		serverOptions.WebPort = CONFIG_DEFAULT_SERVER_PORT
	}
	if c.ServerWebsocketPort > MIN_PORT && c.ServerWebsocketPort < MAX_PORT {
		serverOptions.WebsocketPort = ":" + strconv.Itoa(int(c.ServerWebsocketPort))
	} else {
		serverOptions.WebsocketPort = CONFIG_DEFAULT_SERVER_PORT
	}
	conf.Server = &serverOptions

	//log
	if c.LogLogtofile == "on" {
		logOptions.Logtofile = true
	} else if c.LogLogtofile == "off" {
		logOptions.Logtofile = false
	} else {
		logOptions.Logtofile = CONFIG_DEFAULT_LOG_TO_FILE
	}
	if c.LogLogtostd == "on" {
		logOptions.Logtostd = true
	} else if c.LogLogtostd == "off" {

	} else {
		logOptions.Logtostd = CONFIG_DEFAULT_LOG_TO_STD
	}
	switch c.LogVerbosity {
	case CONFIG_LOG_VERBOSITY_DEBUG:
		logOptions.Verbosity = 0
	case CONFIG_LOG_VERBOSITY_RELEASE:
		logOptions.Verbosity = 1
	default:
		logOptions.Verbosity = CONFIG_DEFAULT_LOG_VERBOSITY
	}
	if c.LogDir != "" {
		logOptions.Dir = c.LogDir
	} else {
		logOptions.Dir = CONFIG_DEFAULT_LOG_DIR
	}
	conf.Log = &logOptions
	//Db
	if c.DbHost != "" {
		mongoOptions.Host = c.DbHost
	} else {
		mongoOptions.Host = mongo.DEFAULT_DB_HOST
	}
	if c.DbPort > 0 {
		mongoOptions.Port = c.DbPort
	} else {
		mongoOptions.Port = mongo.DEFAULT_DB_PORT
	}
	if c.DbName != "" {
		mongoOptions.DbName = c.DbName
	} else {
		mongoOptions.DbName = mongo.DEFAULT_DB_NAME
	}
	mongoOptions.User = c.DbUser
	mongoOptions.Password = c.DbPassword
	mongoOptions.Uri = c.DbUri
	conf.Db = &mongoOptions

	//Debug
	debugOptions.PprofFile = c.DebugPprofFile
	conf.Debug = &debugOptions

	return &conf
}
