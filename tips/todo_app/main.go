package main

import (
	"fmt"

	"github.com/so-heee/go-example/tips/todo_app/app/models"
	"github.com/so-heee/go-example/tips/todo_app/config"
)

func main() {
	fmt.Println(config.Config.DbName)
	fmt.Println(models.Db)
}
