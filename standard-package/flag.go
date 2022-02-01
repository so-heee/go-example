package main

import (
	"flag"
	"fmt"
)

func Flag() {

	// https://pkg.go.dev/flag@go1.17.6

	// サンプル
	// go run main.go -n 20 -m message -x

	// オプションを格納する変数の定義
	var (
		max int
		msg string
		x   bool
	)

	// IntVar 整数のオプション
	flag.IntVar(&max, "n", 32, "処理数の最大値")
	// StringVar 文字列のオプション
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	// BoolVar bool型のオプション コマンドラインに与えられたらtrue なければfalse
	flag.BoolVar(&x, "x", false, "拡張オプション")

	// コマンドラインをパース
	flag.Parse()

	fmt.Println("処理数の最大値 =", max)
	fmt.Println("処理メッセージ =", msg)
	fmt.Println("拡張オプション =", x)
}
