package main

import (
	"fmt"
	. "fmt"
	f "fmt"

	"github.com/so-heee/go-example/udemy/section13/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + "" + Version
}

func Do(s string) (b string) {
	// var d string = s
	b = s
	{
		b := "BBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	//スコープ
	fmt.Println(foo.Max)
	// fmt.Println(foo.min)

	// importの仕方は色々あるがそのまま使用するのがおすすめ
	fmt.Println(foo.ReturnMin())
	f.Println("test")
	Println("test")

	fmt.Println(appName())

	fmt.Println(Do("AAA"))
}
