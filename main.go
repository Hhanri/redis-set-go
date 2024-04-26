package main

import (
	"context"
	"log"
	"time"

	"github.com/Hhanri/redis-set-go/client"
)

func main() {
	go func() {
		server1 := NewServer(Config{})
		log.Fatal(server1.Start())
	}()

	time.Sleep(time.Second)

	client1 := client.New("localhost:5001")

	if err := client1.Set(context.Background(), "foo", "bar"); err != nil {
		log.Fatal(err)
	}

	select {}
}
