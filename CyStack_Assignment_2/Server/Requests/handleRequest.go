package Requests

import (
	"fmt"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Read the data sent by the client
	buf := make([]byte, 1024)
	n, err := r.Body.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received data from client: \n", string(buf[:n])) // ADD THIS LINE

	// Send a response to the client
	_, err = w.Write([]byte("Hi Client, this is the server, I got the data."))
	if err != nil {
		fmt.Println(err)
		return
	}
}
