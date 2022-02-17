package main

import (
	"log"
	"net/http"
)

const DefaultPort = "7893"

// EchoHandler echos back the request as a response
func EchoHandler(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte(request.URL.Path))
}

func main() {

	log.Println("starting backend, listening on port " + DefaultPort)

	http.HandleFunc("/", EchoHandler)
	http.ListenAndServe(":"+DefaultPort, nil)
}
