package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// 構造体からJSONテキストへの変換
type A struct{}

// User.
// 型を変更したい場合はjson:"id,string"
// 表示したくない場合はjson: "-"
// 初期値の場合のみ表示したくない場合はjson: "id,omitempty"
// omitemptyを構造体にする場合は*Aとする
type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Creared time.Time `json:"creared"`
	A       A         `json:"A"`
}

// Marshalのカスタマイズ
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

// Unmarshalのカスタマイズ
func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func Json() {

	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Creared = time.Now()

	// Marshal JSONに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

	fmt.Printf("%T\n", bs)

	u2 := new(User)

	// Unmarshal JSONをデータに変更
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u2)
}
