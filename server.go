package main

import (
    "html/template"
    "log"
    "net/http"
)

var tpl = template.Must(template.ParseFiles("./static/hello.html"))

func hello(w http.ResponseWriter, r *http.Request) {
    tpl.Execute(w, nil)
}

// public serves static assets such as CSS and JavaScript to clients.
//func public() http.Handler {
  //return http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
  //}

func main() {
    http.HandleFunc("/", hello)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
    err := http.ListenAndServe(":8787", nil)
    if err != nil {
        log.Fatal(err)
    }
}
