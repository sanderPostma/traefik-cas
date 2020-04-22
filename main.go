package main

import (
	"traefik-cas/server"
)

// Main
func main() {

	// Setup logger
	log := server.NewDefaultLogger()

	// Build server
	server := server.NewServer()

	log.Info("Starting server on " + server.Config.ListenAddress)
	server.Start()

}
