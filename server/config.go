package server

import (
	"flag"
	"os"
)

var config *Config

type Config struct {
	AuthHost       string
	ListenAddress  string
	UserHeaderName string
}

func NewConfig() *Config {
	c := &Config{}
	flag.StringVar(&c.AuthHost, "cas-base-url", getEnv("CAS_BASE_URL", ""),
		"Sets the CAS server base URL.")
	flag.StringVar(&c.ListenAddress, "listen-address", getEnv("LISTEN_ADDRESS", ":4188"),
		"Sets the listen address. The default is :4188")
	flag.StringVar(&c.UserHeaderName, "user-header-name", getEnv("USER_HEADER_NAME", "X-Forwarded-User"),
		"Configures  the header where the user name value should be set in. The default is X-Forwarded-User")
	return c
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if fallback == "" {
		panic("No value specified for environment variable " + key)
	}
	return fallback
}
