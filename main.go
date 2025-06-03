package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/yuin/goldmark"
)

func main() {
	http.Handle(
		"GET /static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.HandleFunc("GET /", home)
	http.HandleFunc("GET /{slug}", post)

	http.ListenAndServe(":8080", nil)
}

func home(respWriter http.ResponseWriter, req *http.Request) {
	dir, err := os.ReadDir("posts")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	posts := make([]string, len(dir))
	for i := range dir {
		posts[i] = dir[i].Name()
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(respWriter, posts)
}

func post(respWriter http.ResponseWriter, req *http.Request) {
	slug := req.PathValue("slug")
	raw, err := os.ReadFile(fmt.Sprintf("posts/%s.md", slug))
	if err != nil {
		log.Fatal(err)
	}

	if err = goldmark.Convert(raw, respWriter); err != nil {
		log.Fatal(err)
	}
}

type Post struct {
	Title string
	Desc  string
	Url   string
}
