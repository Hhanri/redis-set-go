package protocol

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

const (
	CommandSET = "SET"
)

type Command interface {
}

type SetCommand struct {
	Key string
	Val string
}

func ParseCommand(raw string) (Command, error) {
	rd := resp.NewReader(bytes.NewBufferString(raw))
	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if v.Type() == resp.Array {
			for i, value := range v.Array() {
				switch value.String() {
				case CommandSET:
					return parseSetCommand(v.Array())
				}

				fmt.Printf("  #%d %s, value '%s'\n", i, value.Type(), value)
			}
		}
	}

	return nil, fmt.Errorf("Invalid command received: %s\n", raw)
}

func parseSetCommand(array []resp.Value) (Command, error) {
	if len(array) != 3 {
		return nil, fmt.Errorf("invalid number of variables for SET command")
	}
	cmd := SetCommand{
		Key: array[1].String(),
		Val: array[2].String(),
	}
	return cmd, nil
}
