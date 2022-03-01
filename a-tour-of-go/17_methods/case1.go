package main

import (
	"fmt"
	"math"
)

// Vertex.
type Vertex struct {
	X, Y float64
}

/*

型にメソッドを定義することができる
メソッドは、特別なレシーバ引数を関数に取る
Absメソッドはvという名前のVertex方のレシーバを持つ

*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// レシーバ自身を更新する場合は、変数レシーバではなくポインタレシーバを使用する
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 通常の関数として記述した場合
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 通常の関数として記述する場合でも自身を更新する場合は、ポインタを引数とする
// 呼び出し時にはポインタを渡す必要がある
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// struct型だけでなく、任意の型（type）にもメソッドを宣言できる
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Case1() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())

	// 通常の関数の場合はポインタを渡す必要があるが、メソッドの場合は自動的に解釈される
	v2 := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v2, 5)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	// 逆の場合でも通常の関数の場合は特定の型の変数を渡す必要があるが、メソットの場合は呼び出し側では意識する必要はない
	v3 := Vertex{3, 4}
	fmt.Println(v3.Abs())
	fmt.Println(AbsFunc(v3))

	p2 := &Vertex{4, 3}
	fmt.Println(p2.Abs())
	fmt.Println(AbsFunc(*p2))

}
