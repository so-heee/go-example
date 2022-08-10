package main

import (
	"context"
	"fmt"
	"time"
)

func child(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		fmt.Println("キャンセルされてる")
		return
	}
	fmt.Println("キャンセルされてない")
}

func main() {
	// withCancel()
	done()
}

func withCancel() {
	// 任意のタイミングでキャンセルを通知する
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	cancel()
	child(ctx)
}

func done() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	go func() { fmt.Println("別ゴルーチン") }()
	fmt.Println("STOP")
	<-ctx.Done()
	fmt.Println("そして時は動き出す")
}
