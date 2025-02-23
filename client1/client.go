package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// 2秒のタイムアウトを設定
	ctx, cancel := context.WithTimeout(context.Background(), 2 *time.Second)
	defer cancel()

	// 新しいリクエストを作成
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// コンテキストをリクエストに設定
	req = req.WithContext(ctx)

	// クライアントでリクエストを送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("リクエストエラー:", err)
		return
	}
	defer resp.Body.Close()

	// レスポンスを読み取る
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("サーバーのレスポンス:", string(body))
}