package main

import (
	"log"
	"os"
)

func Log() {

	// https://pkg.go.dev/log@go1.17.6

	// 出力先変更
	log.SetOutput(os.Stdout)
	log.Print("Log\n")
	log.Println("Log2")
	log.Printf("Log%d\n", 3)

	// log.Fatal("Log\n")
	// log.Fatalln("Log2")
	// log.Fatalf("Log%d\n", 3)
	//
	// log.Panic("Log\n")
	// log.Panicln("Log2")
	// log.Panicf("Log%d\n", 3)

	f, err := os.Create("test.log")
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.Println("ファイルに書き込む")

}

func LogFormat() {
	log.SetOutput(os.Stdout)

	// ログのフォーマットを指定する
	// 標準
	log.SetFlags(log.LstdFlags)
	log.Println("A")

	// マイクロ秒を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	// 時刻と行番号（短縮形）
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	log.SetFlags(log.LstdFlags)
	// ログのプレフィックスを追加
	log.SetPrefix("[LOG]")
	log.Println("D")
}

func Logger() {
	// ロガーの生成
	logger := log.New(os.Stdout, "[LOG]", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message")
	log.Println("message")

	// エラーハンドリングで使用
	_, err := os.Open("gsgsde")
	if err != nil {
		logger.Fatalln("Exit", err)
	}

}
