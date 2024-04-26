package main

import "log"

func main() {
	server := NewServer(Config{})
	log.Fatal(server.Start())
}
