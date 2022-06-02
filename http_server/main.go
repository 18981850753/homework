package main

import (
	"cloud/http_server/server"
	"net/http"
)

func main() {
	r := server.NewMux()
	r.HandleFunc("/web", handleRequest)
	r.HandleFunc("/healthz", handleLive)
	http.ListenAndServe(":80", r)
}

func handleRequest(res http.ResponseWriter, req *http.Request) {

}

func handleLive(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("200"))
}
