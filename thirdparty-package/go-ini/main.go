package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// ConfigList.
type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(8080),
		DbName:    cfg.Section("db").Key("name").MustString("example"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("Config.Port = %+v\n", Config.Port)
	fmt.Printf("Config.DbName = %+v\n", Config.DbName)
	fmt.Printf("Config.SQLDriver = %+v\n", Config.SQLDriver)
}
