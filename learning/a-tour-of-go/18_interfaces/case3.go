package main

import "fmt"

/*

空のインターフェースは、任意の型の値を保持できる。 (全ての型は、少なくともゼロ個のメソッドを実装している)
空のインターフェースは、未知の型の値を扱うコードで使用される

*/

func Case3() {
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
