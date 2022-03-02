package main

import "fmt"

// Double.
func Double(i int) {
	i = i * 2
}

// Doublev2.
func Doublev2(i *int) {
	*i = *i * 2
}

// Doublev3.
func Doublev3(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
}

func main() {
	//ポインタ

	var n int = 100

	fmt.Println(n)
	fmt.Println(&n)

	Double(n)

	fmt.Println(n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	// *p = 300
	// fmt.Println(n)
	// n = 200
	// fmt.Println(*p)

	Doublev2(&n)
	fmt.Println(n)

	Doublev2(p)
	fmt.Println(*p)

	//参照型はポインタを使用しなくても参照渡しになる
	var sl []int = []int{1, 2, 3}
	Doublev3(sl)
	fmt.Println(sl)
}
