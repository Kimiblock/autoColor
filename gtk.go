package main

import (
	"os"
	"os/exec"
)

func setGtk(dark bool, conf Config) {
	if ! conf.Gtk.Enable {
		return
	}
	var theme string
	switch dark {
		case true:
			theme = conf.Gtk.Theme.Dark
		case false:
			theme = conf.Gtk.Theme.Light
	}
	logger.Println("Switching GTK theme to", theme)
	cmdSlice := []string{
		"set",
		"org.gnome.desktop.interface",
		"gtk-theme",
		theme,
	}
	cmd := exec.Command("gsettings", cmdSlice...)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		warn.Println("Could not switch GTK theme:", err)
	}
}