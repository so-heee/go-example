package main

import "fmt"

func main() {

	// 型変換
	i := 42
	fmt.Println("int", i)
	f := float64(i)
	fmt.Println("float64", f)
	u := uint(f)
	fmt.Println("uint", u)

	// 定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使える
	// :=の記述はできない
	const Pi = 3.14
	fmt.Print("定数", Pi)
}

/*** Memo ***/

//型一覧
/*
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // uint8 の別名
rune // int32 の別名
     // Unicode のコードポイントを表す
float32 float64
complex64 complex128
*/

// int, uint, uintptr 型は、32-bitのシステムでは32 bitで、64-bitのシステムでは64 bit。
// サイズ、符号なし( unsigned )整数の型を使うための特別な理由がない限り、整数の変数が必要な場合は int を使う
