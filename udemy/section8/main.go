package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func anything(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 100000)
	}
}

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func init() {
	fmt.Println("init")
}

func main() {
	//if
	//条件分岐
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't Know")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)

	//エラーハンドリング
	var s string = "A"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)

	//for
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("Loop")
	}

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for k, v := range arr {
		fmt.Println(k, v)
	}

	sl := []string{"Python", "PHP", "GO"}
	for _, v := range sl {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	//switch
	n := 6
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("i don't know")
	}

	switch n2 := 2; n2 {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("i don't know")
	}

	n3 := 6
	switch {
	case n3 > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n3 > 3 && n < 7:
		fmt.Println("3 < n < 7")
	default:
		fmt.Println("i don't know")
	}

	//型switch
	anything("aaa")
	anything(1)

	var x2 interface{} = 3
	i3 := x2.(int) //最初に定義した型以外にはスイッチできない
	fmt.Println(i3 + 2)

	f, isFloat64 := x2.(float64)
	fmt.Println(f, isFloat64)

	if x2 == nil {
		fmt.Println("None")
	} else if i, isInt := x2.(int); isInt {
		fmt.Println(i, "x2 is Int")
	} else if s, isString := x2.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("i don't know")
	}

	switch x2.(type) { // 該当の値を使用する場合は v := x2.(type)のようにする
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("i don't know")
	}

	//ラベル付きfor
Loop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("処理をしない")
	}
	fmt.Println("END")

Loop2:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop2
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}

	//defer
	TestDefer()

	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))

	//goroutin
	// go sub()

	// for {
	// 	fmt.Println("Main Loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }
}
