package config

import "os"

func Port() string {
	port := os.Getenv("POKEMON_GO_PORT")
	if port == "" {
		return "8080"
	}
	return port
}
