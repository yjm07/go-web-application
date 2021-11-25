package main

import (
	"go-rest-server/basic/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", myapp.NewHttpHandler())
}
