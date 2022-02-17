package main

import (
	"log"
	"net/http"
	"os"
)

const DefaultPort = "7893";

func getServerPort() (string) {
	port := os.Getenv("SERVER_PORT");
	if port != "" {
		return port;
	}

	return DefaultPort;
}

// EchoHandler echos back the request as a response
func EchoHandler(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte(request.URL.Path))
}

func main() {

	log.Println("starting server, listening on port " + getServerPort())

	http.HandleFunc("/", EchoHandler)
	http.ListenAndServe(":" + getServerPort(), nil)
}