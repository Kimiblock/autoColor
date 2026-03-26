package main

type Config struct {
	Qt		QtConf
	Icon		IconConf
	Gtk		GtkConf
}

type IconConf struct {
	Enable	bool
	Theme	ThemeSwitch
}

type GtkConf struct {
	Enable	bool
	Theme	ThemeSwitch
}

type QtConf struct {
	Enable	bool
	Theme	ThemeSwitch
}

type ThemeSwitch struct {
	Light	string
	Dark	string
}