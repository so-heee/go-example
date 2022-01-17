package main

import (
	"fmt"
	"time"
)

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func reciever2(name string, c chan int) {
	for {
		i, ok := <-c
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	// channel
	// 複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造

	var ch1 chan int
	//受信専用
	// var ch2 <-chan int
	//送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	i := <-ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))

	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))

	ch3 <- 1
	fmt.Println("len", len(ch3))
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6

	//チャネルとゴルーチン
	ch4 := make(chan int)
	ch5 := make(chan int)

	go reciever(ch4)
	go reciever(ch5)

	i3 := 0
	for i3 < 100 {
		ch4 <- i3
		ch5 <- i3
		time.Sleep(50 * time.Millisecond)
		i3++
	}

	ch6 := make(chan int, 2)

	// close前にチャネルに値がある場合は受信することができる
	// ch6 <- 1

	close(ch6)

	// ch6 <- 1

	i4, ok := <-ch6
	fmt.Println(i4, ok)

	ch7 := make(chan int, 2)

	go reciever2("1.goroutin", ch7)
	go reciever2("2.goroutin", ch7)
	go reciever2("3.goroutin", ch7)

	i5 := 0
	for i5 < 100 {
		ch7 <- i5
		i5++
	}
	close(ch7)
	time.Sleep(3 * time.Second)

	ch8 := make(chan int, 3)
	ch8 <- 1
	ch8 <- 2
	ch8 <- 3
	close(ch8)
	for k := range ch8 {
		fmt.Println(k)
	}

	ch9 := make(chan int, 3)
	ch10 := make(chan string, 3)

	ch10 <- "A"
	ch9 <- 1
	ch10 <- "B"
	ch9 <- 2

	select {
	case v1 := <-ch9:
		fmt.Println(v1 + 1000)
	case v2 := <-ch10:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch11 := make(chan int)
	ch12 := make(chan int)
	ch13 := make(chan int)

	go func() {
		for {
			i6 := <-ch11
			ch12 <- i6 * 2
		}
	}()

	go func() {
		for {
			i7 := <-ch12
			ch13 <- i7 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch11 <- n:
			n++
		case i8 := <-ch13:
			fmt.Println("recieved", i8)
		default:
			if n > 100 {
				break L
			}
		}
	}
}
