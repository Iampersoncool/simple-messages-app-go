package main

import "os"

const PortEnvKey = "port"

func GetPort() string {
	port := os.Getenv(PortEnvKey)

	if port == "" {
		port = "localhost:3000"
	}

	return port
}
