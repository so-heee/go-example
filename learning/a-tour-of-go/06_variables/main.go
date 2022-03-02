package main

import "fmt"

func main() {
	// 通常の宣言
	var a int
	fmt.Println(a)

	// 初期化"
	var b, c = 1, "abc"
	fmt.Println(b, c)

	// 暗黙的な宣言（代入ではなく宣言であることに注意）"
	d, e, f := 1, "abc", false
	fmt.Println(d, e, f)
}

/*** Memo ***/

// 宣言時に初期値を設定しない場合の値
// 数値型(int,floatなど): 0
// bool型: false
// string型: "" (空文字列( empty string ))
