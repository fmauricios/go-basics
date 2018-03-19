package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var client = &http.Client{Timeout: 10 * time.Second}

	url := "https://jsonplaceholder.typicode.com/posts"

	resp, err := client.Get(url)

	if err != nil {
		panic(err.Error())
	}

	var post []Post

	err = json.NewDecoder(resp.Body).Decode(&post)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(post[0])
}
