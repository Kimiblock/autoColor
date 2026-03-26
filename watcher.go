package main

import (
	"os"
	"os/exec"

	"github.com/godbus/dbus/v5"
)

func watcher(conf Config) {
	conn, err := dbus.SessionBus()
	if err != nil {
		warn.Fatalln("Could not connect to Session Bus:", err)
	}
	portalObj := conn.Object("org.freedesktop.portal.Desktop", "/org/freedesktop/portal/desktop")
	err = conn.AddMatchSignal(
		dbus.WithMatchMember("SettingChanged"),
		dbus.WithMatchObjectPath("/org/freedesktop/portal/desktop"),
		dbus.WithMatchInterface("org.freedesktop.portal.Settings"),
		dbus.WithMatchSender("org.freedesktop.portal.Desktop"),
	)
	if err != nil {
		warn.Fatalln("Could not add match signal:", err)
	}
	sigChan := make(chan *dbus.Signal, 5)
	conn.Signal(sigChan)
	var dark bool
	for sig := range sigChan {
		darkStat := isDark(conn, portalObj)
		if dark == darkStat {
			continue
		}
		logger.Println("Settings changed:", sig.Body)
		dark = darkStat
		switchTheme(darkStat, conf)
	}

}

func switchTheme(dark bool, config Config) {
	qtChange(dark, config)
	setGtk(dark, config)
}

func isDark(conn *dbus.Conn, portalObj dbus.BusObject) bool {
	call := portalObj.Call(
		"org.freedesktop.portal.Settings.ReadOne",
		0,
		"org.freedesktop.appearance",
		"color-scheme",
	)
	if call.Err != nil {
		warn.Fatalln("Could not obtain color-scheme:", call.Err)
	}
	var res uint
	err := call.Store(&res)
	if err != nil {
		warn.Fatalln("Could not obtain color-scheme:", err)
	}
	switch res {
		case 0:
			if os.Getenv("XDG_CURRENT_DESKTOP") == "GNOME" {
				cmdarg := []string{
					"set",
					"org.gnome.desktop.interface",
					"color-scheme",
				}
				cmd := exec.Command("gsettings", cmdarg...)
				cmd.Start()
			}
			return false
		case 1:
			return true
		case 2:
			return false
		default:
			return false
	}
}