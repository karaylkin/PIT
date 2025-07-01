package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", textHandler("Go application is running"))
	http.HandleFunc("/echo", echoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func textHandler(response string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		if _, err := w.Write([]byte(response)); err != nil {
			log.Println("Failed to send response:", err)
		}
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	if _, err := io.Copy(w, r.Body); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error copying request body:", err)
	}
}
