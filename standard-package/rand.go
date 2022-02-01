package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Rand() {

	// https://pkg.go.dev/crypto/rand@go1.17.6

	// 固定値を設定すると何度やっても同じ値になる
	rand.Seed(256)
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	// 現在時刻をシードに使った擬似乱数の生成(このやり方が一般的)
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	// 0-99の間
	fmt.Println(rand.Intn(100))
	// int型
	fmt.Println(rand.Int())
}
