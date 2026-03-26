package main

type Config struct {
	Kvantum		KvConf

}

type IconConf struct {
	Enable	bool
	Theme	ThemeSwitch
}

type GtkConf struct {
	Enable	bool
	Theme	ThemeSwitch
}

type KvConf struct {
	Enable	bool
	Theme	ThemeSwitch
}

type ThemeSwitch struct {
	Light	string
	Dark	string
}