package main

import "os"

func main() {
	// Set value env into global state
	// os.Setenv("API_HOST", "192.168.1.12")
	// os.Setenv("API_HOST", "192.168.0.9")
	os.Setenv("API_HOST", "192.168.0.17")
}
