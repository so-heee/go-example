package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

// ConfigList.
type ConfigList struct {
	Port      int    `required:"true" split_words:"true"`
	DbName    string `required:"true" split_words:"true"`
	SQLDriver string `required:"false" split_words:"true"`
}

var Config ConfigList

func init() {

	if err := envconfig.Process("", &Config); err != nil {
		log.Fatal("error envconfig.Process :", err)
	}
}

func main() {
	fmt.Printf("Config.Port = %+v\n", Config.Port)
	fmt.Printf("Config.DbName = %+v\n", Config.DbName)
	fmt.Printf("Config.SQLDriver = %+v\n", Config.SQLDriver)
}
