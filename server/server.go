package server

import (
	"gopkg.in/cas.v2"
	"net/http"
	"net/url"
)

type Server struct {
	casClient *cas.Client
	mux       *http.ServeMux
}

func NewServer() *Server {
	url, _ := url.Parse("https://cas.dev.sphereon.com/cas")

	return &Server{casClient: cas.NewClient(&cas.Options{
		URL: url,
	})}
}

func (s *Server) Start() {

	s.mux = http.NewServeMux()
	s.mux.Handle("/", s)

	server := &http.Server{
		Addr:    ":4188",
		Handler: NewRequestProcessorHandler(s),
	}

	log.Info(server.ListenAndServe())
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Debugf("Received request from %v", r.URL)

	if !cas.IsAuthenticated(r) {
		log.Debug("Redirecting to CAS login URL")
		cas.RedirectToLogin(w, r)
		return
	}

	if r.URL.Path == "/logout" {
		log.Debug("Logging out")
		cas.RedirectToLogout(w, r)
		return
	}

	username := cas.Username(r)
	log.Debugf("Setting header %s", username)
	w.Header().Set("X-Forwarded-User", username)
	w.WriteHeader(200)
}
