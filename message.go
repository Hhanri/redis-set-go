package main

import "github.com/Hhanri/redis-set-go/protocol"

type Message struct {
	cmd  protocol.Command
	peer *Peer
}
