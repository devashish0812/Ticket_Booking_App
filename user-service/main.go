package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from Go!")
    })

    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}
