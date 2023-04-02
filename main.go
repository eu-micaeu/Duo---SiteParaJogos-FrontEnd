package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// obtém o nome da máquina em que o servidor está sendo executado
	hostname, _ := os.Hostname()

	// exibe uma mensagem informando o nome da máquina e a porta em que o servidor está rodando
	fmt.Printf("Servidor rodando na porta 8080 da máquina %s...\n", hostname)

	// configura os manipuladores de arquivos estáticos
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	// configura os manipuladores de página HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		log.Printf("Solicitação recebida do endereço IP %s para %s", ip, r.URL)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/jogodacobrinha.html", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		log.Printf("Solicitação recebida do endereço IP %s para %s", ip, r.URL)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		http.ServeFile(w, r, "jogodacobrinha.html")
	})

	http.HandleFunc("/jogodaforca.html", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		log.Printf("Solicitação recebida do endereço IP %s para %s", ip, r.URL)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		http.ServeFile(w, r, "jogodaforca.html")
	})

	http.HandleFunc("/jogodamemoria.html", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		log.Printf("Solicitação recebida do endereço IP %s para %s", ip, r.URL)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		http.ServeFile(w, r, "jogodamemoria.html")
	})

	http.HandleFunc("/jogodavelha.html", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		log.Printf("Solicitação recebida do endereço IP %s para %s", ip, r.URL)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		http.ServeFile(w, r, "jogodavelha.html")
	})

	http.HandleFunc("/jogodoclique.html", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		log.Printf("Solicitação recebida do endereço IP %s para %s", ip, r.URL)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		http.ServeFile(w, r, "jogodoclique.html")
	})

	// inicia o servidor
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar servidor: ", err)
	}
}
