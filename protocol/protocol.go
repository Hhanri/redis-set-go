package protocol

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

const (
	CommandSET   = "SET"
	CommandGET   = "GET"
	CommandHello = "hello"
)

type Command interface {
}

type SetCommand struct {
	Key string
	Val []byte
}

type GetCommand struct {
	Key string
}

type HelloCommand struct {
	Value string
}

func parseCommand(r io.Reader, onCommand func(Command), onDone func()) error {
	rd := resp.NewReader(r)
	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			onDone()
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if v.Type() == resp.Array {
			for i, value := range v.Array() {

				var cmd Command
				var err error

				switch value.String() {
				case CommandSET:
					cmd, err = parseSetCommand(v.Array())
				case CommandGET:
					cmd, err = parseGetCommand(v.Array())
				case CommandHello:
					cmd, err = parseHelloCommand(v.Array())
				}

				if err != nil {
					return err
				}

				onCommand(cmd)

				fmt.Printf("  #%d %s, value '%s'\n", i, value.Type(), value)
			}
		}
	}
	return nil
}

func HandleCommand(r io.Reader, onCommand func(Command), onDone func()) error {
	return parseCommand(r, onCommand, onDone)
}

func parseSetCommand(array []resp.Value) (Command, error) {
	if len(array) != 3 {
		return nil, fmt.Errorf("invalid number of variables for SET command")
	}
	cmd := SetCommand{
		Key: array[1].String(),
		Val: array[2].Bytes(),
	}
	return cmd, nil
}

func parseGetCommand(array []resp.Value) (Command, error) {
	if len(array) != 2 {
		return nil, fmt.Errorf("invalid number of variables for GET command")
	}
	cmd := GetCommand{
		Key: array[1].String(),
	}
	return cmd, nil
}

func parseHelloCommand(array []resp.Value) (Command, error) {
	return HelloCommand{
		Value: array[1].String(),
	}, nil
}

func RespWriteMap(m map[string]string) []byte {
	buff := &bytes.Buffer{}
	buff.WriteString("%" + fmt.Sprintf("%d\r\n", len(m)))
	rw := resp.NewWriter(buff)
	for k, v := range m {
		rw.WriteString(k)
		rw.WriteString(":" + v)
	}
	return buff.Bytes()
}
