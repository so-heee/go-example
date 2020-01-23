package calc

import "time"

// Sum は引数を合算して返却する
func Sum(a, b int) int {
	return a + b
}

// SumSleep は3秒待ったのち、引数を合算して返却する
func SumSleep(a, b int) int {
	time.Sleep(3 * time.Second) // 3秒休む
	return a + b
}
