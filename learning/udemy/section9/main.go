package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	//スライス
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)
	fmt.Println(sl5[0])
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])
	fmt.Println(sl5[:])
	fmt.Println(sl5[len(sl5)-1])
	fmt.Println(sl5[1 : len(sl5)-1])

	// append len cap
	sl6 := []int{100, 200}
	fmt.Println(sl6)

	sl6 = append(sl6, 300)
	fmt.Println(sl6)

	sl6 = append(sl6, 400, 500, 600)
	fmt.Println(sl6)

	sl7 := make([]int, 5)
	fmt.Println(sl7)
	fmt.Println(len(sl7))
	fmt.Println(cap(sl7))

	sl8 := make([]int, 5, 10)
	fmt.Println(sl8)
	fmt.Println(len(sl8))
	fmt.Println(cap(sl8))

	sl8 = append(sl7, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(len(sl8))
	fmt.Println(cap(sl8))

	sl9 := []int{100, 200}
	sl10 := sl9
	sl10[0] = 1000
	fmt.Println(sl9)
	//参照型は同じアドレスと参照するため元の値も変更される

	sl11 := []int{1, 2, 3, 4, 5}
	sl12 := make([]int, 5, 10)
	fmt.Println(sl12)
	n := copy(sl12, sl11) // nにはコピーされた要素数が格納される
	fmt.Println(n, sl12)

	//スライスfor
	sl13 := []string{"A", "B", "C"}
	fmt.Println(sl13)

	for k, v := range sl13 {
		fmt.Println(k, v, sl13[k])
	}

	//可変長引数
	fmt.Println(Sum(1, 2, 3, 4))
	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Sum())
	sl14 := []int{1, 2, 3}
	fmt.Println(Sum(sl14...))
}
