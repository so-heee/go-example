package main

import (
	"fmt"
	"math"
)

func main() {
	// 大文字で始まる関数はパッケージ外からアクセス可能
	// 小文字で始まる関数はパッケージ外からアクセス不可
	fmt.Println(math.Pi)
}
