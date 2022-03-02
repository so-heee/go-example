package main

import (
	"fmt"
)

func main() {
	hello()
	counting()
}

func hello() {
	// defer ステートメントは、 defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させる
	// defer へ渡した関数の引数は、すぐに評価されるが、その関数自体は呼び出し元の関数がreturnするまで実行されない
	defer fmt.Println("world")
	fmt.Println("hello")
}

func counting() {
	fmt.Println("counting")

	// defer へ渡した関数が複数ある場合、その呼び出しはスタック( stack )され
	// LIFO(last-in-first-out:後入先出) の順番で実行される
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
