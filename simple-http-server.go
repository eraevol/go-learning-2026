package main

import (
	"net/http"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().String()
	w.Write([]byte(currentTime))
}
func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8080", nil)
}
