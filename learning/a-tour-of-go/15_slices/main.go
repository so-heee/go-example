package main

import "fmt"

func main() {
	// 配列は固定長である代わりに、スライスは可変長である
	a := [6]int{2, 3, 5, 7, 11, 13}

	// 1番目から3番目の要素を取得（3, 5, 7）
	var s []int = a[1:4]
	fmt.Println(s)

	// 新規配列を作成し、それを参照するスライスを作成
	fmt.Println([]bool{true, true, false})

	// 全て0から9番目まで取得
	var b [10]int
	fmt.Println(b[0:10])
	fmt.Println(b[:10])
	fmt.Println(b[0:])
	fmt.Println(b[:])

	c := []int{2, 3, 5, 7, 11, 13}
	// スライスの長さ(length)は、それに含まれる要素の数
	fmt.Println(len(c))
	// スライスの容量(capacity)は、スライスの最初の要素から数えて、元となる配列の要素数
	fmt.Println(cap(c))
	c = c[:4]
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
	c = c[2:]
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	// スライスのゼロ値は nil 、nil 長さと容量は 0
	var d []int
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	// 動的サイズの配列を作成
	// a len=5 cap=5 [0 0 0 0 0]
	e := make([]int, 5)
	fmt.Println(e)
	// b len=0 cap=5 []
	k := make([]int, 0, 5)
	fmt.Println(k)
	// c len=2 cap=5 [0 0]
	f := e[:2]
	fmt.Println(f)
	// d len=3 cap=3 [0 0 0]
	g := f[2:5]
	fmt.Println(g)
	// e len=5 cap=5 [0 0 0 0 0]
	h := []int{0, 0, 0, 0, 0}
	fmt.Println(h)
	// f len=2 cap=5 [0 0]
	i := h[:2]
	fmt.Println(i)
	// g len=3 cap=3 [0 0 0]
	j := i[2:5]
	fmt.Println(j)

	// 多次元のスライス
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)

	// 新しい要素の追加
	l := make([]int, 2)
	fmt.Println(l)
	l = append(l, 3, 4, 5)
	fmt.Println(l)

	// Rangeの一つ目（i）はインデックス、二つ目はインデックスの場所の要素のコピー
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Println(i, v)
	}
	// インデックスを破棄する場合は_
	for _, v := range pow {
		fmt.Println(v)
	}
	// 値を破棄する場合は, vを省略
	for i := range pow {
		fmt.Println(i)
	}

}
