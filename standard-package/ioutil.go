package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Ioutil() {

	// https://pkg.go.dev/io/ioutil@go1.17.6

	// 入力全体を読み込む
	f, _ := os.Open("foo.txt")
	// 巨大な入力データでは利用しない方がいい
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))

	// ファイルの書き込み
	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}
}
