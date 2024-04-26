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
}

func New(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (c *Client) Set(ctx context.Context, key string, val string) error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	wr := resp.NewWriter(&buff)
	wr.WriteArray(
		[]resp.Value{
			resp.StringValue(protocol.CommandSET),
			resp.StringValue(key),
			resp.StringValue(val),
		},
	)

	_, err = conn.Write(buff.Bytes())
	return err
}
