package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// net/client

func Client() {
	// Get
	res, _ := http.Get("https://example.com")
	fmt.Println(res.StatusCode)
	fmt.Println(res.Proto)
	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])

	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	// Post
	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())

	res, err := http.PostForm("http://example.com/", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body2, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body2))
}

func CustomGet() {
	//応用
	//ヘッダーをつけたり、クエリをつけたり
	//Parse 正しいURLか確認
	base, _ := url.Parse("https://example.com/")

	//クエリ の確認 URLの後につく
	reference, _ := url.Parse("index/lists?id=1")

	//ResolveReference
	//クエリをくっつけたURLを生成する。
	//相対パスから絶対URLに変換する。
	//baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	//GET ver
	//リクエストを作成 nil部はPOST時のみ設定（バイトを入れる）
	//まだリクエストはしていない。structを作っただけ。
	req, _ := http.NewRequest("GET", endpoint, nil)

	//requestにheaderをつける。cash情報など
	req.Header.Add("Content-Type", `application/json"`)

	//URLのクエリを確認
	q := req.URL.Query()

	//クエリを追加
	q.Add("name", "test")

	//クエリを表示
	fmt.Println(q)

	//&など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
	fmt.Println(q.Encode())

	//encodeしてからURLに戻す
	//日本語などを変換する
	req.URL.RawQuery = q.Encode()

	//実際にアクセスする
	//クライアントを作る
	var client *http.Client = &http.Client{}

	//結果 アクセスする
	resp, _ := client.Do(req)

	//読み込み
	body, _ := ioutil.ReadAll(resp.Body)

	//出力
	fmt.Println(string(body))
}

//応用
type Account struct {
	ID       string
	PassWord string
}

func CustomPost() {
	//応用
	//ヘッダーをつけたり、クエリをつけたり
	//Parse 正しいURLか確認
	base, _ := url.Parse("https://example.com/")

	//クエリ の確認 URLの後につく
	reference, _ := url.Parse("index/lists?id=1")

	//Postの時のデータ
	AccountData := &Account{ID: "ABC-DEF", PassWord: "testtest"}
	data, _ := json.Marshal(AccountData)

	//ResolveReference
	//クエリをくっつけたURLを生成する。
	//相対パスから絶対URLに変換する。
	//baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	//POST ver
	//bytes.NewBuffer([]byte("password")でリクエストの領域を作成
	//POSTの場合は、Bodyにデータを入れる。例えばパスワード。見られたらダメな情報はbodyに
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))

	//requestにheaderをつける。cash情報など
	req.Header.Add("Content-Type", `application/json"`)

	//URLのクエリを確認
	q := req.URL.Query()

	//クエリを追加
	q.Add("name", "test")

	//クエリを表示
	fmt.Println(q)

	//&など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
	fmt.Println(q.Encode())

	//encodeしてからURLに戻す
	//日本語などを変換する
	req.URL.RawQuery = q.Encode()

	//実際にアクセスする
	//クライアントを作る
	var client *http.Client = &http.Client{}

	//結果 アクセスする
	resp, _ := client.Do(req)

	//読み込み
	body, _ := ioutil.ReadAll(resp.Body)

	//出力
	fmt.Println(string(body))
}
