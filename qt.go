package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func kvChange(dark bool, conf Config) {
	var theme string
	switch dark {
		case true:
			theme = conf.Qt.Theme.Dark
		case false:
			theme = conf.Qt.Theme.Light
	}
	logger.Println("Switching Qt theme to", theme)
	os.Mkdir(filepath.Join(confDir, "qt6ct"), 0700)
	file, err := os.OpenFile(
		filepath.Join(confDir, "qt6ct", "qt6ct.conf"),
		os.O_RDONLY,
		0700,
	)
	if err != nil {
		if os.IsNotExist(err) {
			builder := strings.Builder{}
			_, err := builder.WriteString("[Appearance]\n")
			if err != nil {
				panic(err)
			}
			_, err = builder.WriteString("style=" + theme + "\n")
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(
				filepath.Join(confDir, "qt6ct", "qt6ct.conf"),
				[]byte(builder.String()),
				0700,
			)
			if err != nil {
				warn.Println("Could not switch Qt theme:", err)
			}
			return
		}
		warn.Println("Could not switch Qt theme:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	builder := strings.Builder{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "style=") {
			_, err := builder.WriteString("style=" + theme + "\n")
			if err != nil {
				panic(err)
			}
			continue
		}
		_, err := builder.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	file.Close()
	err = os.WriteFile(
		filepath.Join(confDir, "qt6ct", "qt6ct.conf.new"),
		[]byte(builder.String()),
		0700,
	)
	if err != nil {
		warn.Println("Could not switch Qt theme:", err)
		return
	}
	err = os.Rename(
		filepath.Join(confDir, "qt6ct", "qt6ct.conf.new"),
		filepath.Join(confDir, "qt6ct", "qt6ct.conf"),
	)
	if err != nil {
		warn.Println("Could not switch Qt theme:", err)
		return
	}
}