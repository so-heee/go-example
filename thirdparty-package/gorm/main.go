package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User ユーザーモデル
type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

// Users ユーザーモデル（複数形）
type Users []User

// NewSQLHandler データベース接続
func NewSQLHandler() *gorm.DB {
	conn, err := gorm.Open("mysql", "root:root@tcp(localhost)/sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return conn
}

func main() {
	db := NewSQLHandler()
	defer db.Close()
	Init(db)
	Create(db)
}

// Init デーブルの削除 マイグレーション
func Init(db *gorm.DB) {
	fmt.Println("process Init")
	// `User`モデルのテーブルを削除します
	db.DropTable(&User{})
	// `User`モデルのテーブルをマイグレーションします
	db.AutoMigrate(&User{})
}

// Create 新規作成
func Create(db *gorm.DB) {
	fmt.Println("process Create")
	user := User{
		Name:      "test",
		CreatedAt: time.Now(),
	}
	db.Create(&user)
	fmt.Println(user)
}
