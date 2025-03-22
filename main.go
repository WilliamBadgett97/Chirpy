package main

import (
	"net/http"
)
func main() {
	server := &Server{
		server: http.NewServeMux(),
		addr: ":8080",
	}
	http.ListenAndServe(server.addr,server.server)
}