package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8069"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Порт: %s", port)
	})

	log.Printf("Сервер работает на порту %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
