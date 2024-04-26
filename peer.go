package main

import (
	"net"

	"github.com/Hhanri/redis-set-go/protocol"
)

type Peer struct {
	conn  net.Conn
	msgCh chan<- Message
}

func NewPeer(conn net.Conn, msgCh chan<- Message) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
	}
}

func (p *Peer) Send(msg []byte) (int, error) {
	return p.conn.Write(msg)
}

func (p *Peer) readLoop() error {
	for {
		err := protocol.HandleCommand(p.conn, func(cmd protocol.Command) {
			p.msgCh <- Message{
				cmd:  cmd,
				peer: p,
			}
		})
		if err != nil {
			return err
		}
	}
}
