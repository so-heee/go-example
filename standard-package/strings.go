package main

import (
	"fmt"
	"strings"
)

func Strings() {

	// https://pkg.go.dev/strings@go1.17.6

	// 文字列を結合する
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2)

	// 文字列に含まれる部分文字列を検索する
	i1 := strings.Index("ABCDE", "A")
	i2 := strings.Index("ABCDE", "D")
	fmt.Println(i1, i2)
	i3 := strings.LastIndex("ABCDEABCDE", "BC")
	fmt.Println(i3)
	i4 := strings.IndexAny("ABC", "ABC")
	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5)
	b1 := strings.HasPrefix("ABC", "A")
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2)
	b3 := strings.Contains("ABC", "B")
	b4 := strings.ContainsAny("ABCDE", "BD")
	fmt.Println(b3, b4)
	b5 := strings.Count("ABCABC", "B")
	b6 := strings.Count("ABCABC", "")
	fmt.Println(b5, b6)

	// 文字列を繰り返し結合
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4)

	// 文字列の置換
	s5 := strings.Replace("AAAAAA", "A", "B", 1)
	s6 := strings.Replace("AAAAAA", "A", "B", -1)
	s7 := strings.ReplaceAll("AAAAAA", "A", "B")
	fmt.Println(s5, s6, s7)

	// 文字列の分割
	s8 := strings.Split("A,B,C,D,E", ",")
	fmt.Println(s8)
	s9 := strings.SplitAfter("A,B,C,D,E", ",")
	fmt.Println(s9)
	s10 := strings.SplitN("A,B,C,D,E", ",", 2)
	fmt.Println(s10)
	s11 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
	fmt.Println(s11)

	// 大文字、小文字の変換
	s12 := strings.ToLower("ABC")
	s13 := strings.ToLower("E")
	s14 := strings.ToUpper("abc")
	s15 := strings.ToUpper("e")
	fmt.Println(s12, s13, s14, s15)

	// 文字列から空白を取り除く
	s16 := strings.TrimSpace("    -    ABC    -    ")
	// 全角
	s17 := strings.TrimSpace("　　　ABC　　　")
	fmt.Println(s16, s17)

	// 文字列からスペースで区切られたスペースを取り出す
	s18 := strings.Fields("a b c")
	fmt.Println(s18)
	fmt.Println(s18[1])
}
