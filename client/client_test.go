package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestNewClient(t *testing.T) {
	client1, err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if err := client1.Set(
			context.Background(),
			fmt.Sprintf("foo_%d", i),
			fmt.Sprintf("bar_%d", i),
		); err != nil {
			log.Fatal(err)
		}

		b, err := client1.Get(
			context.Background(),
			fmt.Sprintf("foo_%d", i),
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}

	client1.Close()
}

func TestMultipleClients(t *testing.T) {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(n int) {

			client, err := New("localhost:5001")
			defer func() {
				client.Close()
				wg.Done()
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

}
