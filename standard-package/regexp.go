package main

import (
	"fmt"
	"regexp"
)

func Regexp() {

	// https://pkg.go.dev/regexp@go1.17.6

	// 正規表現の基本 regexp.MatchString
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)

	// Compile （予めコンパイルしておくことで、上記の書き方と違って都度コンパイルされない）
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match)

	// MustCompile
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match)

	// regexp.MustCompile("\\d")
	// regexp.MustCompile(`\d`)
}

func RegexpFlag() {

	/*
		フラグ一覧

		i 大文字小文字を区別しない
		m マルチラインモード（^と$が文頭、文末に加えて行頭、行末にマッチ）
		s .が\nにマッチ
		U 最小マッチへの変換（x*はx*?へ、x+はx+?へ）
	*/

	re3 := regexp.MustCompile(`(?i)abc`)
	match := re3.MatchString("ABC")
	fmt.Println(match)

	/*
		幅をもたない正規表現のパターン
		パターン一覧

		^ 文頭
		$ 文末
		\A 文頭
		\z 文末
		\b ASCIIによるワード境界
		\B 非ASCIIによるワード境界
	*/

	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match)

	match = re4.MatchString("  ABC  ")
	fmt.Println(match)

	/*
		繰り返しを表す正規表現
		繰り返しのパターン

		x* 0回以上繰り返すx（最大マッチ）
		x+ 1回以上繰り返すx（最大マッチ）
		x? 0回以上1回以下繰り返すx
		x{n,m} n回以上m回以下繰り返すx（最大マッチ）
		x{n,} n回以上繰り返すx（最大マッチ）
		x{n} n回繰り返すx（最大マッチ）
		x*? 0回以上繰り返すx（最小マッチ）
		x+? 1回以上繰り返すx（最小マッチ）
		x?? 0回以上1回以下繰り返すx（0回優先）
		x{n,m}? n回以上m回以下繰り返すx（最小マッチ）
		x{n,}? n回以上繰り返すx（最小マッチ）
		x{n}? n回繰り返すx（最小マッチ）
	*/

	re5 := regexp.MustCompile(`a+b*`)
	fmt.Println(re5.MatchString("ab"))
	fmt.Println(re5.MatchString("a"))
	fmt.Println(re5.MatchString("aaaaaabbbbbbbb"))
	fmt.Println(re5.MatchString("b"))

	// 正規表現の文字クラス
	re6 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re6.MatchString("Y"))

	re7 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re7.MatchString("ABC"))
	fmt.Println(re7.MatchString("abcdefg"))

	// 英数字と_以外の文字にマッチ
	re8 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re8.MatchString("ABC"))
	fmt.Println(re8.MatchString("あ"))

	/*
		正規表現のグループ

		(正規表現) グループ（順序によるキャプチャ）
		(?:正規表現) グループ（キャプチャされない）
		(?:P<name>正規表現) 名前付きグループ
	*/
	re9 := regexp.MustCompile((`(abc|ABC)|(xyz|XYZ)`))
	fmt.Println(re9.MatchString("abcxyz"))
	fmt.Println(re9.MatchString("ABCXYZ"))
	fmt.Println(re9.MatchString("ABCxyz"))
	fmt.Println(re9.MatchString("ABCabc"))
}
