package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)


// LoggingTransport構造体は、http.RoundTripper関数を実装する構造体
type LoggingTransport struct {
	Transport http.RoundTripper
}

// RoundTripメソッドをオーバーライドしてリクエストごとにログを出力
func (lt *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	log.Printf("リクエスト: %s %s", req.Method, req.URL)

	// Transportがnilの場合、http.DefaultTransportを使用
	if lt.Transport == nil {
		lt.Transport = http.DefaultTransport
	}

	// 実際のリクエストを送信
	resp, err := lt.Transport.RoundTrip(req)
	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		return nil, err
	}

	// レスポンスのステータスのログを出力
	log.Printf("レスポンス: %s %s [%d] - %v", req.Method, req.URL, resp.StatusCode, time.Since(start))

	return resp, nil
}

func main() {
	// http.DefaultTransportをカスタムトランスポートとして利用
	client := &http.Client{
		Transport: &LoggingTransport{Transport: http.DefaultTransport},
		Timeout: 5 * time.Second, // タイムアウトを設定
	}

	// APIにGETリクエストを送信
	url := "https://jsonplaceholder.typicode.com/posts/1"
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("GETリクエスト失敗: %v", err)
	}
	defer resp.Body.Close()

	// レスポンスボディを読み込む
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ボディレスポンス読み取り失敗: %v", err)
	}

	// 結果を出力
	fmt.Println("レスポンスボディ:")
	fmt.Println(string(body))
}