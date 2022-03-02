package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// ConfigList.
type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	name := os.Getenv("DB_NAME")
	driver := os.Getenv("SQL_DRIVER")

	Config = ConfigList{
		Port:      port,
		DbName:    name,
		SQLDriver: driver,
	}

}

func main() {
	fmt.Printf("Config.Port = %+v\n", Config.Port)
	fmt.Printf("Config.DbName = %+v\n", Config.DbName)
	fmt.Printf("Config.SQLDriver = %+v\n", Config.SQLDriver)
}
