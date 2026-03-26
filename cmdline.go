package main

import "os"

func cmdline(args []string) Config {
	var ret Config
	var skip int
	var dryRun bool
	for idx := range args {
		if skip > 0 {
			skip--
			continue
		}
		switch args[idx] {
			case "--config":
				skip++
				if len(args) < idx + 1 {
					warn.Fatalln("--config requires an argument")
				}
				ret = readConf(args[idx + 1])
			case "--dry-run":
				dryRun = true
			default:
				warn.Println("Unknown argument:", args[idx])
		}
	}
	if dryRun {
		logger.Println(ret)
		os.Exit(0)
	}
	return ret
}