package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/jogodacobrinha", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "jogodacobrinha.html")
	})

	http.HandleFunc("/jogodaforca", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "jogodaforca.html")
	})

	http.HandleFunc("/jogodamemoria", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "jogodamemoria.html")
	})

	http.HandleFunc("/jogodavelha", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "jogodavelha.html")
	})

	http.HandleFunc("/jogodoclique", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "jogodoclique.html")
	})

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
