package main

import (
	"fmt"
	"math"
)

/*

変数レシーバではなくポインタレシーバを使うときの切り分け

- メソッドがレシーバが指す先の変数を変更する場合
- メソッドの呼び出し毎に変数のコピーを避けるため（レシーバが大きな構造体である場合に効果的）

一般的には値レシーバ、またはポインタレシーバを混合させるべきではない

*/

// Vertex.
type Vertex2 struct {
	X, Y float64
}

/*

型にメソッドを定義することができる
メソッドは、特別なレシーバ引数を関数に取る
Absメソッドはvという名前のVertex方のレシーバを持つ

*/
func (v *Vertex2) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// レシーバ自身を更新する場合は、変数レシーバではなくポインタレシーバを使用する
func (v *Vertex2) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Case2() {
	v := &Vertex2{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs2())
	v.Scale2(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs2())
}
