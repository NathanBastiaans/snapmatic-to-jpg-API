package main

import (
	"os"
	"snapmatic-to-jpg/src/web"
)

// main starts the web server
func main() {
	if err := web.Start(); err != nil {
		os.Exit(1)
	}
}
