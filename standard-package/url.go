package main

import (
	"fmt"
	"net/url"
)

// net/url

func Url() {
	// URL解析
	u, _ := url.Parse("http://example.com/search?a=1&b=2#top")
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)
	fmt.Println(u.Query())

	// URLを生成
	url := &url.URL{}
	url.Scheme = "https:"
	url.Host = "google.com"
	q := url.Query()
	q.Set("q", "Go言語")
	url.RawQuery = q.Encode()
	fmt.Println(url)
}
