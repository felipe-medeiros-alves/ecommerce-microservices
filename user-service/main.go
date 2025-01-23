package main

import (
    "fmt"
    "net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "User Profile")
}

func main() {
    http.HandleFunc("/users", userHandler)
    http.ListenAndServe(":8002", nil)
}
