package main

import (
	"go-rest-server/users/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", myapp.NewHandler())
}
