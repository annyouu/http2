package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Person構造体を定義 (JSONデータをマッピング)
type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
}

func main() {
	// GETリクエストを送信
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatal("GETリクエストエラー:", err)
	}
	// 関数終了時にレスポンスボディを閉じる
	defer resp.Body.Close()

	// HTTPステータスコードを確認
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("リクエスト失敗: ステータスコード %d", resp.StatusCode)
	}

	// JSONデータをデコードしてPerson構造体にマッピング
	var p Person
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&p); err != nil {
		log.Fatal("JSONデコードエラー:", err)
	}

	// デコードしたPersonの情報を表示
	fmt.Println("取得したデータ:")
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("名前: %s\n", p.Name)
	fmt.Printf("ユーザー名: %s\n", p.Username)
	fmt.Printf("メール: %s\n", p.Email)
}


// // JSONからGo構造体への変換 (デコード)
// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	ID int `json:"id"`
// 	Name string `json:"name"`
// 	Username string `json:"username"`
// 	Email string `json:"email"`
// }

// func main() {
// 	jsonString := `{"id":2,"name":"佐藤花子","username":"sato", "email":"sato@example.com"}`

// 	var person Person
// 	json.Unmarshal([]byte(jsonString), &person)

// 	fmt.Println(person)
// }

