package main

import (
	"fmt"
	"log"
	"os"
)

// https://pkg.go.dev/os@go1.17.6

// Exit.
func Exit() {
	// os.Exit(1)
	// fmt.Println("Start")

	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}

// Fatal.
func Fatal() {
	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

// Args.
func Args() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	fmt.Printf("length=%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}

}

// File.
func File() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

}

// FileCreate.
func FileCreate() {
	f, _ := os.Create("foo.txt")
	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang"), 6)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yaah")
}

// FileRead.
func FileRead() {
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bs := make([]byte, 128)
	n, _ := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

	bs2 := make([]byte, 128)

	nn, _ := f.ReadAt(bs2, 10)
	fmt.Println(nn)
	fmt.Println(string(bs2))
}

// FileOpen.
func FileOpen() {
	f, err := os.OpenFile("foo.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(n)
	fmt.Println(string(bs))

}
