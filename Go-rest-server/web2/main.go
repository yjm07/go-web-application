package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// read file from client
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	// make directory to save the file
	dir_name := "./uploads"
	os.MkdirAll(dir_name, 0777)
	filepath := fmt.Sprintf("%s/%s", dir_name, header.Filename)
	file, err := os.Create(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	defer file.Close()

	// copy file to the directory
	_, err_ := io.Copy(file, uploadFile)
	if err_ != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err_)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.HandleFunc("/upload", uploadHandler)

	http.ListenAndServe(":8000", nil)
}
