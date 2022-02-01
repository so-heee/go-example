package main

import (
	"fmt"
	"os"
)

func Fmt() {

	// https://pkg.go.dev/fmt

	// 標準
	fmt.Print("Hello")
	// Println 各々の文字列は半角スペースで区切られ、文字列の最後に改行が挿入される
	fmt.Println("Hello", "Hello")
	fmt.Println("Hello", "Hello")

	// Printf フォーマットの指定
	fmt.Printf("%s\n", "Hello")
	fmt.Printf("%v\n", "Hello")

	// Sprint 出力ではなくフォーマットした結果を文字列で返す
	s := fmt.Sprint("Hello")
	s1 := fmt.Sprintf("%v\n", "Hello")
	s2 := fmt.Sprintln("Hello")

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)

	// Fprint 書き込み先指定
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test1.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint")
}

/*

	Goは書式指定子 %... のことを verb と表記しています。

   %v

   %vはさまざまな型の値を柔軟に出力することができる。
   基本的なデータ型に対して%vは下記の書式指定子として振舞う。interface型として本来の型が不定であるデータの内容を出力させる場合に有用。
   スライス、マップ、配列、構造体についても有用。


   すべての型に使えるverb

   論理値(bool)   %t
   符号付き整数(int, int8など) %d
   符号なし整数(uint, uint8など)   %d
   浮動小数点数(float64など)   %g
   複素数(complex128など)   %g
   文字列(string) %s
   チャネル(chan)  %p
   ポインタ(pointer)   %p


   %+v  %vと同じだが、構造体の場合にフィールド名を出力する。
   %#v  値のGoの文法での表現を出力する。
   %T   値の型のGoの文法での表現を出力する。
   %%   %そのものを出力したい場合に使う。


   論理値に使えるverb

   %t trueかfalse


   整数に使えるverb

   %d  10進数で出力する
   %+d 10進数で出力し、正の整数でも符号を付与する
   %(数値)d  10進数で出力し、(数値)で指定した桁数だけ左を半角スペースで埋める
   %-(数値)d 10進数で出力し、(数値)で指定した桁数だけ右を半角スペースで埋める
   %0(数値)d 10進数で出力し、(数値)で指定した桁数だけ左を0で埋める
   %b  2進数で出力する
   %o  8進数で出力する
   %#o 0付き8進数で出力する
   %x  16進数で出力する(a-fは小文字)
   %#x 0x付き16進数で出力する(a-fは小文字)
   %X  16進数で出力する(A-Fは大文字)
   %#X 0x付き16進数で出力する(A-Fは大文字)
   %U  Unicodeコードポイントに対応する文字で出力する


   浮動小数点数・複素数に使えるverb

   %f  実数表現で出力する
   %F  実数表現で出力する(%fと同じ)
   %e  仮数と指数表現で出力する(eは小文字)
   %E  仮数と指数表現で出力する(Eは大文字)
   %g  指数部が大きい場合は%e、それ以外は%fで出力する
   %G  指数部が大きい場合は%E、それ以外は%Fで出力する


   文字列([]byteも同じ)に使えるverb

   %s  そのままの書式で出力する
   %(数値)s  (数値)で指定した桁数だけ左を半角スペースで埋める
   %e  Goの文法上のエスケープをした文字列で出力する(ダブルクォート付与)
   %q
   %x  16進数表現で出力する(a-fは小文字)
   %X  16進数表現で出力する(A-Fは大文字)
*/
