package main

import (
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	t.Run("Without port env var", func(t *testing.T) {
		got := GetPort()
		want := "localhost:3000"

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})

	t.Run("With port env var", func(t *testing.T) {
		os.Setenv(PortEnvKey, "localhost:8080")

		got := GetPort()
		want := "localhost:8080"

		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}
