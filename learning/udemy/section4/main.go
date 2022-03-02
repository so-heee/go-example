package main

import (
	"fmt"
	"strconv"
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

	//byte(uint8)型
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

	//配列型
	//後からサイズの変更ができない
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	//中身から要素数を決める
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	arr2[2] = "C"
	fmt.Println(arr2)
	fmt.Println(len(arr1))

	//interface型
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	// 計算などはできない
	// fmt.Println(x + 1)
	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	//型変換
	var i3 int = 1
	fl642 := float64(i3)

	fmt.Println(fl642)
	fmt.Printf("i3 = %T\n", i3)
	fmt.Printf("fl642 = %T\n", fl642)

	i4 := int(fl642)
	fmt.Printf("i4 = %T\n", i4)

	fl643 := 10.5
	//小数点以下を切り捨て
	i5 := int(fl643)
	fmt.Printf("i5 = %T\n", i5)
	fmt.Println(i5)

	var s2 string = "100"
	fmt.Printf("s2 = %T\n", s2)

	i6, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i6)
	fmt.Printf("i6 = %T\n", i6)

	var i7 int = 200
	s3 := strconv.Itoa(i7)
	fmt.Println(s3)
	fmt.Printf("s3 = %T\n", s3)

	var s4 string = "Hello World"
	b := []byte(s4)
	fmt.Println(b)

	s5 := string(b)
	fmt.Println(s5)
}
