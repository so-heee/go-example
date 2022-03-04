package main

import "fmt"

// グローバル変数
var i8 = 100

func main() {

	// 明示的な宣言
	var i int
	var f float64
	fmt.Println(i)
	fmt.Println(f)

	// 明示的な宣言＋初期化
	var i2 int = 100
	fmt.Println(i2)

	// 暗黙的な宣言
	var i3 = 100
	fmt.Println(i3)

	// var の省略
	i4 := 100
	fmt.Println(i4)

	// まとめて宣言
	var i5, s = 100, "abc"
	var (
		i6 = 100
		s2 = "abc"
	)
	i7, s3 := 100, "abc"
	fmt.Println(i5, i6, i7, s, s2, s3)

	fmt.Println(I)

}

/*

宣言時に初期値を設定しない場合の値
- 数値型(int,floatなど): 0
- bool型: false
- string型: "" (空文字列( empty string ))
- その他: nil



*/
