package main

import (
	"fmt"
	"net/http"
)

type loggingTransport struct {
	Transport http.RoundTripper
}

func (lt *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// リクエストの詳細をログに出力
	fmt.Printf("Request: %s %s\n", req.Method, req.URL)

	// 実際のリクエスト処理を親のTransportに委託
	if lt.Transport == nil {
		lt.Transport = http.DefaultTransport
	}
	return lt.Transport.RoundTrip(req)
}

func main() {
	client := &http.Client {
		Transport: &loggingTransport{},
	}

	// HTTPリクエストを送信
	resp, err := client.Get("https://www.example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}