package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")			// Hello, Worldってアクセスした人に返すよ！
}

func main() {
	http.HandleFunc("/", handler) 			// http://localhost:8080/にアクセスしてきた人はhandlerを実行するよ！
	http.ListenAndServe(":8080", nil)// サーバーを起動するよ！
}