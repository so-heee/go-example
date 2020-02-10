package main

import "fmt"

// Vertex 構造体を宣言
type Vertex struct {
	X int
	Y int
}

// 初期化の種類
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// 構造体を宣言
	v := Vertex{1, 2}
	// ドット( . )でフィールドにアクセス
	v.X = 4
	fmt.Println(v.X)
	// ポインタを通してアクセス
	p := &v
	p.X = 5
	fmt.Println(v.X)

	fmt.Println(v1, p, v2, v3)
}
