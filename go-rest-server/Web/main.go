package main

import (
	"go-rest-server/Web/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", myapp.NewHttpHandler())
}
