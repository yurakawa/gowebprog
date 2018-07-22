package main

import (
	"net/http"
	"html/template"
	"time"
	"math/rand"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html")
	t.ExecuteTemplate(w, "layout", "")
}

func process2(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	num := rand.Intn(10)
	if num > 7 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else if (num > 3){
		t, _ = template.ParseFiles("layout.html", "blue_hello.html")
	} else {
		t, _ = template.ParseFiles("layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")


}
func main() {
	server := http.Server {
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	server.ListenAndServe()
}


