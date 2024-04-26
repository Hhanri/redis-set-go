package main

import (
	"fmt"
	"net"

	"github.com/Hhanri/redis-set-go/protocol"
	"golang.org/x/exp/slog"
)

const defaultListenAddr = ":5001"

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	peers map[*Peer]bool
	ln    net.Listener

	addPeerCh chan *Peer
	quitCh    chan struct{}
	msgCh     chan Message

	kv *KV
}

func NewServer(cfg Config) *Server {
	if cfg.ListenAddr == "" {
		cfg.ListenAddr = defaultListenAddr
	}
	return &Server{
		Config:    cfg,
		peers:     make(map[*Peer]bool),
		addPeerCh: make(chan *Peer),
		quitCh:    make(chan struct{}),
		msgCh:     make(chan Message),
		kv:        NewKV(),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln
	go s.loop()

	slog.Info("server running", "listenAddr", s.ListenAddr)
	return s.acceptLoop()
}

func (s *Server) handleMessage(msg Message) error {
	cmd, err := protocol.ParseCommand(string(msg.data))
	if err != nil {
		return err
	}

	switch v := cmd.(type) {
	case protocol.SetCommand:
		return s.kv.Set(v.Key, v.Val)
	case protocol.GetCommand:
		val, ok := s.kv.Get(v.Key)
		if !ok {
			return fmt.Errorf("Key not found")
		}
		_, err := msg.peer.Send(val)
		if err != nil {
			slog.Error("Peer send error", "err", err)
			return err
		}
	}
	return nil
}

func (s *Server) loop() {
	for {
		select {
		case msg := <-s.msgCh:
			if err := s.handleMessage(msg); err != nil {
				slog.Error("raw message error", "err", err)
			}
		case <-s.quitCh:
			return
		case peer := <-s.addPeerCh:
			s.peers[peer] = true
		}
	}
}

func (s *Server) acceptLoop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("accept error", "err", err)
			continue
		}

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	peer := NewPeer(conn, s.msgCh)
	s.addPeerCh <- peer

	slog.Info("New peer connected", "remoteAddr", conn.RemoteAddr())
	if err := peer.readLoop(); err != nil {
		slog.Error("peer read error", "err", err, "remoteAddr", conn.RemoteAddr())
	}
}
