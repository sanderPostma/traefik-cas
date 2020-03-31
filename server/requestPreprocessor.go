package server

import (
	"net/http"
	"net/url"
)

type RequestProcessor struct {
	server *Server
}

func NewRequestProcessorHandler(server *Server) http.Handler {
	return &RequestProcessor{server: server}
}

func (rp *RequestProcessor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Method = r.Header.Get("X-Forwarded-Method")
	r.Host = r.Header.Get("X-Forwarded-Host")
	r.URL, _ = url.Parse(r.Header.Get("X-Forwarded-Uri"))

	rp.server.casClient.Handle(rp.server.mux).ServeHTTP(w, r)
}
