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
}
