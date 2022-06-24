package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/so-heee/go-example/thirdparty-package/ent/ent"
	"github.com/so-heee/go-example/thirdparty-package/ent/start"
)

func main() {
	entOptions := []ent.Option{}

	// 発行されるSQLをロギングするなら
	entOptions = append(entOptions, ent.Debug())

	client, err := ent.Open("mysql", "root:root@tcp(localhost)/sample?charset=utf8&parseTime=True&loc=Local", entOptions...)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// 発行しているSQLを確認する場合
	// if err := client.Schema.WriteTo(ctx, os.Stdout, migrate.WithForeignKeys(false)); err != nil {
	// 	log.Fatalf("failed printing schema changes: %v", err)
	// }

	_, err = start.CreateUser(ctx, client)
	if err != nil {
		log.Fatalf("failed create user: %v", err)
	}

	_, err = start.QueryUser(ctx, client)
	if err != nil {
		log.Fatalf("failed query user: %v", err)
	}

	log.Print("ent sample done.")
}
