package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/packer/common/json"
)

func TestHandleGet(t *testing.T) {
	// 各テストケースは独立して実行されるため、それぞれにテスト用のWebサーバを立ち上げる必要がある
	mux := http.NewServeMux()               // テストを実行するマルチプレクサを生成
	mux.HandleFunc("/post/", handleRequest) // テスト対象のハンドラを追加

	writer := httptest.NewRecorder()                     // 返されたHTTPレスポンスを取得
	request, _ := http.NewRequest("GET", "/post/1", nil) // テストしたいハンドラ宛のリクエストを作成
	mux.ServeHTTP(writer, request)                       // テスト対象のハンドラにリクエストを送信

	if writer.Code != 200 { //ResponseRecoderにより結果をチェック
		t.Errorf("Response code is %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.ID != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
