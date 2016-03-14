package common

import (
	"fmt"
	"os"
	"os/exec"
)

func Daemon(daemon bool) {
	args := os.Args[1:]
	isDaemon := false
	for i := 0; i < len(args); i++ {
		if args[i] == "daemon" {
			isDaemon = true
			break
		}
	}

	if daemon && !isDaemon {
		args = append(args, "daemon")

		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("run as daemon mode [PID]", cmd.Process.Pid)
		os.Exit(0)
	}
}
