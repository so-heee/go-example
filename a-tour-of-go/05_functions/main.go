package main

import "fmt"

func main() {
	// 引数の型が同じ場合はまとめる
	fmt.Println(add(42, 13))

	// 関数の二つ以上の引数が同じ型の場合は省略可能
	fmt.Println(add2(42, 13))

	// 複数の戻り値を返す場合
	fmt.Println(swap("hello", "world"))

	// 戻り値に名前をつける場合（可読性は低そう）
	fmt.Println(split(17))

	// 関数値（関数の戻り値として利用した場合）
	a := func(x, y int) int {
		return x + y
	}
	// 関数値（関数の引数として利用した場合）
	fmt.Println(compute(a))

	// クロージャ
	b, c := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(b(i), c(i*-2))
	}
}

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func compute(fn func(int, int) int) int {
	// 引数として受け取ったa関数に引数（3, 4）をわたして呼び出し
	return fn(3, 4)
}

// クロージャ
func adder() func(int) int {
	// adderはint型を引数としてとり、int型を返す関数を返す
	// adderで返す下記の関数からはsumを参照でき、adderを呼び出して新しい関数を返すたびにそれぞれの関数に別々のsumが紐づく
	sum := 0
	//closure関数
	return func(x int) int {
		sum += x
		return sum
	}
}
