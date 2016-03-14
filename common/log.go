package common

import (
	"github.com/nprog/SkyEye/log"
)

type LogOptions struct {
	Logtofile bool
	Dir       string
	Logtostd  bool
	Verbosity log.Level
}

func SetLogSettings(lo *LogOptions) {
	log.SetLogOptions(lo.Logtofile, lo.Logtostd, lo.Verbosity)
	log.SetLogDir(lo.Dir)
}
