package main

import (
	"flag"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", defaultListenAddr, "listen address of the redis-set-go server")

	server := NewServer(
		Config{
			ListenAddr: *listenAddr,
		},
	)

	log.Fatal(server.Start())
}
