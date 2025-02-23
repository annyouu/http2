// サーバー側のコード

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, req *http.Request) {
	// リクエストのコンテキストを取得
	ctx := req.Context()
	fmt.Println("リクエスト開始")

	select {
	case <- time.After(5 * time.Second):
		fmt.Fprintln(w, "処理が完了しました")
		fmt.Println("処理完了")
		//クライアントがリクエストをキャンセルしたとき
	case <- ctx.Done():
		fmt.Println("リクエストがキャンセルされました:", ctx.Err())
		http.Error(w, "リクエストがキャンセルされました:", http.StatusRequestTimeout)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("サーバーを起動します...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}