package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Auth service started")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Auth Service")
    })

    http.ListenAndServe(":8000", nil)
}
