package main

import "fmt"

func main() {
	// ポインタは値のメモリアドレスのこと

	// 変数に値を設定
	i, j := 42, 2701

	// & オペレータで値を引き出す
	p := &i
	// * オペレータでポインタの指す先の変数を参照
	// ポインタpを通してiから値を読みだす
	fmt.Println(*p)
	// ポインタpを通してiへ値を代入する
	*p = 21
	// iの値は21になる
	fmt.Println(i)

	p = &j
	// ポインタはjの値を参照
	fmt.Println(*p)
	*p = *p / 37
	// jの値は73になる
	fmt.Println(j)
}
