package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/Hhanri/redis-set-go/client"
)

func TestServerWithMultiClients(t *testing.T) {

	server := NewServer(Config{})

	go func() {
		log.Fatal(server.Start())
	}()

	time.Sleep(time.Second)

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(n int) {

			client, err := client.New("localhost:5001")
			defer func() {
				wg.Done()
				client.Close()
			}()

			if err != nil {
				log.Fatal(err)
			}

			for ii := 0; ii < 10; ii++ {

				key := fmt.Sprintf("client_%d_foo_%d", n, ii)
				if err := client.Set(
					context.Background(),
					key,
					fmt.Sprintf("client_%d_bar_%d", n, ii),
				); err != nil {
					log.Fatal(err)
				}

				b, err := client.Get(
					context.Background(),
					key,
				)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Client %d go this => %s\n", i, b)
			}

		}(i)

	}

	wg.Wait()

	time.Sleep(time.Second)

	if len(server.peers) != 0 {
		t.Errorf("Expected 0 peers but got %d peers", len(server.peers))
	}
}
