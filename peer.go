package main

import (
	"net"
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
	buff := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buff)
		if err != nil {
			return err
		}

		msgBuff := make([]byte, n)
		copy(msgBuff, buff[:n])
		p.msgCh <- Message{
			data: msgBuff,
			peer: p,
		}
	}
}
