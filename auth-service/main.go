package main

import (
    "fmt"
    "net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Login successful")
}

func main() {
    http.HandleFunc("/login", loginHandler)
    http.ListenAndServe(":8000", nil)
}
