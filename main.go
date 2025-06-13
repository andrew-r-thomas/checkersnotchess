package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
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
	md := goldmark.New(goldmark.WithExtensions(meta.Meta))
	posts := getPosts(md)
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(respWriter, TemplData{Posts: posts, Content: template.HTML("hello there")})
}

func post(respWriter http.ResponseWriter, req *http.Request) {
	md := goldmark.New(goldmark.WithExtensions(meta.Meta))
	posts := getPosts(md)

	slug := req.PathValue("slug")
	raw, err := os.ReadFile(fmt.Sprintf("posts/%s.md", slug))
	if err != nil {
		log.Fatal(err)
	}

	ctx := parser.NewContext()
	var buf bytes.Buffer
	if err = md.Convert(raw, &buf, parser.WithContext(ctx)); err != nil {
		log.Fatal(err)
	}
	metaData := meta.Get(ctx)
	title := metaData["title"].(string)
	desc := metaData["description"].(string)

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(
		respWriter,
		TemplData{
			Posts:   posts,
			Content: template.HTML(buf.String()),
			Slug:    slug,
			Title:   title,
			Desc:    desc,
		},
	)
}

func getPosts(md goldmark.Markdown) []Post {
	dir, err := os.ReadDir("posts")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	ctx := parser.NewContext()
	posts := make([]Post, len(dir))
	for i := range dir {
		src, err := os.ReadFile(fmt.Sprintf("posts/%s", dir[i].Name()))
		if err != nil {
			log.Fatal(err)
		}
		var buf bytes.Buffer
		if err := md.Convert(src, &buf, parser.WithContext(ctx)); err != nil {
			log.Fatal(err)
		}
		metaData := meta.Get(ctx)
		title := metaData["title"].(string)
		desc := metaData["description"].(string)
		posts[i] = Post{
			Title: title,
			Desc:  desc,
			Url:   fmt.Sprintf("%s", strings.TrimSuffix(dir[i].Name(), ".md")),
		}
	}

	return posts
}

type Post struct {
	Title string
	Desc  string
	Url   string
}
type TemplData struct {
	Posts   []Post
	Content template.HTML
	Title   string
	Desc    string
	Slug    string
}
