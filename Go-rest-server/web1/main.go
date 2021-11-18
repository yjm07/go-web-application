package main

import (
	"go-rest-server/web1/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", myapp.NewHttpHandler())
}
