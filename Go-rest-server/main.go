package main

import (
	"go-rest-server/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", myapp.NewHttpHandler())
}
