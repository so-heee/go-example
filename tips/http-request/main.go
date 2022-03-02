package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestParams struct {
	Name string `json:"name"`
}

type ResponseParams struct {
	Login   string `json:"login"`
	ID      int    `json:"id"`
	HtmlURL string `json:"html_url"`
}

func main() {
	params := new(RequestParams)
	params.Name = "sample"
	jsonParams, _ := json.Marshal(params)
	fmt.Printf("[+] %s\n", string(jsonParams))

	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/so-heee", bytes.NewReader(jsonParams))
	if err != nil {
		fmt.Printf("Error NewRequest:%T message: %v", err, err)
	}

	fmt.Println(params)
	req.Header.Set("Content-Type", "application/json")
	// API Tokenが必要な場合はセット
	// req.Header.Set("","")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error Request code:%d message: %v", resp.StatusCode, err)
	}
	defer resp.Body.Close()

	var response ResponseParams
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}

	fmt.Println(response.ID)
}
