package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
)

func main() {
	//新しいコンテキストを作成(タイムアウトを指定)
	ctx,cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel() // 関数終了時にcancel()を呼び出してリソースを解放

	// 新しいリクエストを作成し、コンテキストを設定
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx) // 新しいコンテキストを送信するリクエストに送信

	// リクエストを送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// レスポンス処理
	fmt.Println("Status Code:", resp.StatusCode)

	// レスポンスのボディを読み取る例
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response Body", string(body))
}