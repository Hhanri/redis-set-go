package main

import (
	"fmt"
	"net"

	"github.com/Hhanri/redis-set-go/protocol"
)

type Peer struct {
	conn  net.Conn
	msgCh chan<- Message
	delCh chan<- *Peer
}

func NewPeer(conn net.Conn, msgCh chan<- Message, delCh chan<- *Peer) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
		delCh: delCh,
	}
}

func (p *Peer) Send(msg []byte) (int, error) {
	return p.conn.Write(msg)
}

func (p *Peer) readLoop() error {
	for {
		err := protocol.HandleCommand(
			p.conn,
			func(cmd protocol.Command) {
				p.msgCh <- Message{
					cmd:  cmd,
					peer: p,
				}
			},
			func() {
				p.Disconnect()
			},
		)
		if err != nil {
			return err
		}
	}
}

func (p *Peer) Disconnect() {
	fmt.Println("Disconnecting peer")
	p.delCh <- p
}
