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

func main() {
    http.HandleFunc("/", hello)

    err := http.ListenAndServe(":8787", nil)
    if err != nil {
        log.Fatal(err)
    }
}
