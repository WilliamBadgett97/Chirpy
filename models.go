package main

import "net/http"


type Server struct {
	server *http.ServeMux
	addr string
}

