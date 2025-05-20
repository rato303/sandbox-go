package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, ECS!")
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", handler)
    log.Printf("Listening on :%s...", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}