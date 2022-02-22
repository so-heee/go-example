package main

import (
	"fmt"
	"sync"
	"time"
)

// ミューテックスによる同期処理
var st struct{ A, B, C int }

// ミューテックスを保持するパッケージ変数
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	// ロック
	mutex.Lock()
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)
	// アンロック
	mutex.Unlock()
}

func Sync() {
	mutex = new(sync.Mutex)
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	for {

	}
}

// ゴルーチンの終了を待ち受ける

func SyncWait() {
	// sync.WaitGroupを生成
	wg := new(sync.WaitGroup)
	// 待ち受けるゴルーチンの数は3
	wg.Add(3)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2st Goroutine")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3st Goroutine")
		}
		wg.Done()
	}()

	// ゴルーチンの完了を待ち受ける
	// Doneが3つ完了するまで待つ
	wg.Wait()
	fmt.Println("Done")
}
