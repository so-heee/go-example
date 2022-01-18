package main

import "fmt"

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

// T.
type T struct {
	User User
}

// T2.
type T2 struct {
	User
}

// User.
type User struct {
	Name string
	Age  int
}

type Users []*User

// NewUser.
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

// UpdateUser.
func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

// UpdateUser2.
func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

// SayName.
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	// user5 := User{50, "user5"}
	// fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)

	//メソッド
	user9 := User{Name: "user9"}
	user9.SayName()

	user9.SetName("A")
	user9.SayName()

	user9.SetName2("A")
	user9.SayName()

	user10 := &User{Name: "user10"}
	user10.SetName2("B")
	user10.SayName()

	t := T{User: User{Name: "usert", Age: 99}}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)

	//フィールド名を省略した場合は直接参照できる
	t2 := T2{User{Name: "usert2", Age: 100}}
	fmt.Println(t2.Name)

	t.User.SetName2("C")
	fmt.Println(t)

	//コンストラクタ
	user11 := NewUser("user11", 11)
	fmt.Println(user11)
	fmt.Println(*user11)

	user12 := User{Name: "user12", Age: 12}
	user13 := User{Name: "user13", Age: 13}
	user14 := User{Name: "user14", Age: 14}
	user15 := User{Name: "user15", Age: 15}

	users := Users{}
	users = append(users, &user12)
	users = append(users, &user13)
	users = append(users, &user14, &user15)

	fmt.Println(users)
	for _, u := range users {
		fmt.Println(*u)
	}

	users2 := make(Users, 0)
	users2 = append(users2, &user12, &user13)

	for _, u := range users2 {
		fmt.Println(*u)
	}

	//マップ
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}
	fmt.Println(m)
	m2 := map[User]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}
	fmt.Println(m2)
	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "user3"}
	fmt.Println(m3)

	for _, v := range m3 {
		fmt.Println(v)
	}

	//独自型
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("mi = %T\n", mi)

	// i := 100
	// 独自型のため計算できない
	// fmt.Println(mi + i)

	mi.Print()
}
