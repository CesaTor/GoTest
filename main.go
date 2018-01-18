package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	"os"
)

func RootHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "HOME")
}

func MemeHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	tpl, err := template.ParseGlob("assets/templates/*")
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(res, "meme.gohtml", nil)
	fmt.Fprintln(res, "<h1>This is my meme</h1>")
}

func main() {
	http.HandleFunc("/", MemeHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./assets/img"))))

	log.Fatal(http.ListenAndServe(":80", nil))
}
