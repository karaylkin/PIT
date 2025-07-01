package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "6980"

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Порт: %s", port)
}

func main() {
	http.HandleFunc("/", handler)

	log.Printf("Сервер работает на порту %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
