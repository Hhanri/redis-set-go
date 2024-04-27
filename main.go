package main

import (
	"flag"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", defaultListenAddr, "listen address of the redis-set-go server")

<<<<<<< Updated upstream
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

		b, err := client1.Get(
			context.Background(),
			fmt.Sprintf("foo_%d", i),
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(b)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(server1.kv.data)
=======
	server := NewServer(
		Config{
			ListenAddr: *listenAddr,
		},
	)
>>>>>>> Stashed changes

	log.Fatal(server.Start())
}
