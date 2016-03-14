package main

import (
	"github.com/Terry-Mao/goconf"
	"github.com/nprog/SkyEye/common"
	"github.com/nprog/SkyEye/log"
)

//Config 读取配置
type Config struct {
	Server *ServerOptions
	Log    *common.LogOptions
	Debug  *common.DebugOptions
}

//ConfigTemp 文件格式
type ConfigTemp struct {
	//server
	ServerServerAddr string `goconf:"server:server_addr"`
	ServerDaemon     string `goconf:"server:daemon"`

	//log
	LogLogtofile string `goconf:"log:logtofile"`
	LogDir       string `goconf:"log:dir"`
	LogLogtostd  string `goconf:"log:logtostd"`
	LogVerbosity string `goconf:"log:verbosity"`
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
		debugOptions  common.DebugOptions
	)

	//Server
	if c.ServerServerAddr != "" {
		serverOptions.ServerAddr = c.ServerServerAddr
	} else {
		serverOptions.ServerAddr = CONFIG_DEFAULT_SERVER_ADDR
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

	//Debug
	debugOptions.PprofFile = c.DebugPprofFile
	conf.Debug = &debugOptions

	return &conf
}
