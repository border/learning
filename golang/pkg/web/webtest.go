package main

import (
    "fmt"
    "http"
    "log"
    "runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    log.Println("WebServer Start Listen 8080")
    runtime.GOMAXPROCS(4)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
