package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	TILESDB = tilesDB()
	fmt.Println("Mosaic server started")
	server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	// bodyのメモリ上の最大容量は10MB
	r.ParseMultipartForm(10485760)

	// アップロードされたファイルとタイルサイズを取得
	file, _, _ := r.FormFile("image")
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))

	// アップロードされたターゲット画像をデコード
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()

	// newimage := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.X, bounds.Max.Y))

	// タイル画像データベースを複製
	db := cloneTilesDB()

	// 独立して処理できるように画像を分割して処理を分散(ファン・アウト)
	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	c2 := cut(original, &db, tileSize, bounds.Min.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := cut(original, &db, tileSize, bounds.Min.X/2, bounds.Min.Y/2, bounds.Max.X, bounds.Max.Y)

	// 画像をつなげて処理結果をまとめる
	c := combine(bounds, c1, c2, c3, c4)

	// JPEGにエンコードし、base64文字列に変更してブラウザに送信
	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic":   <-c
		"duration": fmt.Sprintf("%v", t1.Sub(t0)),
	}
	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)

}
