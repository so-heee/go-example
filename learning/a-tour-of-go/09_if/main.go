package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("通常のif文", sqrt(2), sqrt(-4))
	fmt.Println("評価前ステートメント", pow(2, 2, 10), pow(2, 3, 20))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// 評価前に簡単なステートメントを書ける（スコープはif内）
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
