package main

import "fmt"

func main() {
	sum := 0
	// 初期化と後処理ステートメントは省略可能
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// golangにはwhile構文が存在しないためforで実装する
	sum2 := 1
	for sum2 < 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}
