package main

import "os"

const PortEnvKey = "PORT"

func GetPort() string {
	port := os.Getenv(PortEnvKey)

	if port == "" {
		port = "localhost:3000"
	}

	return port
}
