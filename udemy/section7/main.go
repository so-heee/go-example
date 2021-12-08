package main

import "fmt"

//関数
func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
}

//関数を返す関数
func ReturnFUnc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

//関数を引数にとる関数
func CallFunction(f func()) {
	f()
}

//クロージャー
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

//ジェネレーター
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)
	i22, _ := Div(9, 4)
	fmt.Println(i22)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()

	//無名関数
	f := func(x, y int) int {
		return x + y
	}
	i5 := f(1, 2)
	fmt.Println(i5)

	i6 := func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(i6)

	f2 := ReturnFUnc()
	f2()

	CallFunction(func() {
		fmt.Println("i'm a function2")
	})

	f3 := Later()
	fmt.Println(f3("Hello"))
	fmt.Println(f3("My"))
	fmt.Println(f3("name"))
	fmt.Println(f3("is"))
	fmt.Println(f3("Golang"))

	f4 := integers()
	fmt.Println(f4())
	fmt.Println(f4())
	fmt.Println(f4())
	fmt.Println(f4())
	fmt.Println(f4())

	f5 := integers()
	fmt.Println(f5())
	fmt.Println(f5())
	fmt.Println(f5())
	fmt.Println(f5())
}
