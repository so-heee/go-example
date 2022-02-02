package main

import (
	"bufio"
	"fmt"
	"os"
)

func Bufio() {

	// https://pkg.go.dev/bufio@go1.17.6

	// 標準入力を行単位で読み込む

	// 標準入力をソースにしたスキャナの生成
	scanner := bufio.NewScanner(os.Stdin)

	// 入力のスキャンが成功する限り繰り返すループ
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// スキャン中にエラーが発生した場合
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー", err)
	}
}
