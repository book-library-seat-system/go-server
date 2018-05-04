package main

import (
	"os"

	"github.com/book-library-seat-system/go-server/server"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8888"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	ser := server.NewServer()
	ser.Run(":" + port)
}
