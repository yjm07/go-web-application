package main

import "net/http"

func uploadHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.HandleFunc("/upload", uploadHandler)

	http.ListenAndServe(":8000", nil)
}
