package main

import "fmt"

func main() {
	// 配列の長さは、型の一部分になるため、配列のサイズは変更できない
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 通常の変数と同じで:=で宣言可能
	b := [2]string{"Test", "Test2"}
	b[0] = "Hello"
	b[1] = "World"
	fmt.Println(b[0], b[1])
	fmt.Println(b)
}
