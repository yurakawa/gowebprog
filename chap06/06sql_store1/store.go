package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

// Db represents database handle
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.ID, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.ID)
	return
}

func main() {
	post := Post{Content: "Hello World", Author: "Sau Sheong"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.ID)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Mondel"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete()

	posts, _ = Posts(10)
	fmt.Println(posts)
}
