package common

import (
	"github.com/nprog/SkyEye/log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

type DebugOptions struct {
	PprofFile string
}

func Pprof(pprofFile *string) {
	f, err := os.Create(*pprofFile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("Ctrl+C to quit.")
		pprof.StopCPUProfile()
		os.Exit(1)
	}()
}
