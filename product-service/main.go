package main

import (
    "fmt"
    "net/http"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Product List")
}

func main() {
    http.HandleFunc("/products", productHandler)
    http.ListenAndServe(":8001", nil)
}
