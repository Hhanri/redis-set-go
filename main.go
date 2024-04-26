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

	for i := 0; i < 10; i++ {
		if err := client1.Set(
			context.Background(),
			fmt.Sprintf("foo_%d", i),
			fmt.Sprintf("bar_%d", i),
		); err != nil {
			log.Fatal(err)
		}
	}

	time.Sleep(time.Second * 3)
	fmt.Println(server1.kv.data)

}
