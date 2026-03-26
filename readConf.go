package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

func readConf (path string) Config {
	file, err := os.OpenFile(path, os.O_RDONLY, 0700)
	if err != nil {
		warn.Fatalln("Could not open configuration:", err)
	}
	defer file.Close()
	decoder := toml.NewDecoder(file)
	var ret Config
	md, err := decoder.Decode(&ret)
	if err != nil {
		warn.Fatalln("Could not decode configuration:", err)
	}
	if len(md.Undecoded()) > 0 {
		warn.Println("Could not decode part of the configuration:", md.Undecoded())
	}
	return ret
}