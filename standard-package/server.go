package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// net/http

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "Hello World111")
}

func Server() {
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil) //基本的にnil指定で問題ない（デフォルトのServeMux型ハンドラ）
}
