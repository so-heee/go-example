package main

import (
	"fmt"

	"github.com/so-heee/golang-example/testing/pkg/benchmark"
	"github.com/so-heee/golang-example/testing/pkg/calc"
	"github.com/so-heee/golang-example/testing/pkg/examples"
	"github.com/so-heee/golang-example/testing/pkg/testmain"
	"github.com/so-heee/golang-example/testing/pkg/unorderedoutput"
)

func main() {
	fmt.Println(calc.Sum(1, 2))
	fmt.Println(examples.Hello())
	unorderedoutput.Unordered()
	fmt.Println(benchmark.Cat("aaa", "bbb", "ccc"))
	fmt.Println(benchmark.Buf("aaa", "bbb", "ccc"))
	fmt.Println(testmain.HelloA())
	fmt.Println(testmain.HelloB())
}
