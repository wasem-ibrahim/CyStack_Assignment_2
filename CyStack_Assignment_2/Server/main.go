package main

import (
	"net/http"
	tp "servergo/Requests"
)

func main() {

	http.HandleFunc("/", tp.HandleRequest)
	http.ListenAndServe(":8080", nil)
}
