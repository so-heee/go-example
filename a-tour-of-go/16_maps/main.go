package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var q = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// mapに渡すトップレベルの型が単純な型名である場合は、その型名を省略可能
var k = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func WordCount(s string) map[string]int {
	result := map[string]int{}
	for _, word := range strings.Fields(s) {
		result[word]++
	}
	return result
}

func main() {
	m = make(map[string]Vertex)
	// mapはキーと値を関連付けます
	// mapの初期値はnil
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(q)
	fmt.Println(k)

	l := make(map[string]int)
	l["Answer"] = 42
	fmt.Println("The value:", l["Answer"])
	l["Answer"] = 48
	fmt.Println("The value:", l["Answer"])
	delete(l, "Answer")
	fmt.Println("The value:", l["Answer"])
	// キーに対する要素が存在するかどうか
	v, ok := l["Answer"]
	fmt.Println("THe value:", v, "Present?", ok)

	wc.Test(WordCount)
}
