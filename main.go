package main

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stdout, "autoColor ", 0)
	warn = log.New(os.Stderr, "autoColor warning ", 0)
	confDir, _ = os.UserConfigDir()
)

func main () {
	if len(confDir) == 0 {
		warn.Fatalln("Could not determine configuration directory")
	}
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	switch len(args) {
		case 0, 1:
			logger.Println("autoColor started with argument:", args)
		default:
			logger.Println("autoColor started with arguments:", args)
	}

	conf := cmdline(args)
	watcher(conf)

}