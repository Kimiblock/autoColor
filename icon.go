package main

import (
	"os"
	"os/exec"
)

func setIcon(dark bool, conf Config) {
	if ! conf.Icon.Enable {
		return
	}
	var theme string
	switch dark {
		case true:
			theme = conf.Icon.Theme.Dark
		case false:
			theme = conf.Icon.Theme.Light
	}
	logger.Println("Switching icon theme to", theme)
	cmdSlice := []string{
		"set",
		"org.gnome.desktop.interface",
		"icon-theme",
		theme,
	}
	cmd := exec.Command("gsettings", cmdSlice...)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		warn.Println("Could not switch icon theme:", err)
	}
}