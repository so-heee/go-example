package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	models "github.com/so-heee/go-example/thirdparty-package/sqlboiler/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "user:password@tcp(localhost)/sample")
	users, err := models.Users(
		Select("id", "name"),
		Where("id > ?", 2),
		Limit(5),
	).All(ctx, db)
	if err != nil {
		fmt.Println("test")
	}

	fmt.Println("users:", users)
	for _, u := range users {
		fmt.Println(u.Name)
	}
}
