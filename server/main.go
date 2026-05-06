package main

import (
	"log"

	"lampcache/server/server"
)

func main() {
	log.Println("Starting lampache service")
	server.StartAndRunTcpServer()
}
