package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })
    fmt.Println("Servidor rodando na porta 8080...")
    http.ListenAndServe(":8080", nil)
}