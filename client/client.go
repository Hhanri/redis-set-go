package client

import (
	"bytes"
	"context"
	"net"

	"github.com/Hhanri/redis-set-go/protocol"
	"github.com/tidwall/resp"
)

type Client struct {
	addr string
	conn net.Conn
}

func New(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Client{
		addr: addr,
		conn: conn,
	}, nil
}

func (c *Client) Set(ctx context.Context, key string, val string) error {
	var buff bytes.Buffer
	wr := resp.NewWriter(&buff)
	wr.WriteArray(
		[]resp.Value{
			resp.StringValue(protocol.CommandSET),
			resp.StringValue(key),
			resp.StringValue(val),
		},
	)

	_, err := c.conn.Write(buff.Bytes())
	return err
}

func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	var buff bytes.Buffer
	wr := resp.NewWriter(&buff)
	wr.WriteArray(
		[]resp.Value{
			resp.StringValue(protocol.CommandGET),
			resp.StringValue(key),
		},
	)

	_, err := c.conn.Write(buff.Bytes())
	if err != nil {
		return nil, err
	}

	rBuff := make([]byte, 35)
	n, err := c.conn.Read(rBuff)
	return rBuff[:n], err
}
