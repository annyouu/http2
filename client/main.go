package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.Clientを作成
	client := &http.Client{}

	// http.NewRequestでリクエストを作成
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		log.Fatal("リクエスト作成エラー:", err)
	}

	// ヘッダーを追加
	req.Header.Add("IF-None-Match", `W/"wyzzy"`)

	// http.Client.Doを使用してリクエストを送信
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("リクエスト送信エラー:", err)
	}
	defer resp.Body.Close() //関数終了時にレスポンスボディを閉じる

	// ステータスコードを表示
	fmt.Println("ステータスコード:", resp.StatusCode)
}