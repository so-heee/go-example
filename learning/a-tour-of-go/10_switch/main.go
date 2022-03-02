package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 選択された case だけを実行してそれに続く全ての case は実行されない（自動的にbreakされる）
	// 条件有りの場合
	withConditions()
	// 条件無しの場合
	noCondition()
}

func withConditions() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func noCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
