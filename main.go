package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Hhanri/redis-set-go/client"
)

func main() {
	server1 := NewServer(Config{})

	go func() {
		log.Fatal(server1.Start())
	}()

	time.Sleep(time.Second)

	client1 := client.New("localhost:5001")

	if err := client1.Set(context.Background(), "foo", "bar"); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	fmt.Println(server1.kv.data)

}
