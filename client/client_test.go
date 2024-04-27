package client

import (
	"context"
	"fmt"
	"log"
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
}
