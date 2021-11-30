package main

import (
	"fmt"
)

//型

func main() {
	//整数型
	var i int = 100

	/*
		// int8 is the set of all signed 8-bit integers.
		// Range: -128 through 127.
		type int8 int8

		// int16 is the set of all signed 16-bit integers.
		// Range: -32768 through 32767.
		type int16 int16

		// int32 is the set of all signed 32-bit integers.
		// Range: -2147483648 through 2147483647.
		type int32 int32

		// int64 is the set of all signed 64-bit integers.
		// Range: -9223372036854775808 through 9223372036854775807.
		type int64 int64

		intを指定した場合は暗黙的に環境依存によっていずれかの型になる
		同じintでもそれぞれ別の型として認識される
	*/

	fmt.Println(i + 50)

	//型を確認することができる
	var i2 int64 = 200
	fmt.Printf("%T\n", i2)
	//型変換
	fmt.Printf("%T\n", int32(i2))
	fmt.Println(i + int(i2))

	//浮動小数点型
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	//暗黙的な定義の場合はfloat64になる
	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 3.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	// 特殊な値(+Inf, -Inf, NaN)
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	// 使用頻度が低い
	// uint(+の整数)
	// complex(複素数型)

	//論理値型
	var t, f bool = true, false
	fmt.Println(t, f)

	//文字列型
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`
	test
		test
	test
	`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))
}
