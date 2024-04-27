package protocol

import (
	"bytes"
	"io"
	"testing"
)

func TestParseCommad(t *testing.T) {
	msg := "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"

	var cmd any

	err := parseCommand(
		bytes.NewBufferString(msg),
		func(_cmd Command) {
			cmd = _cmd
		},
		func() {},
	)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	expected := SetCommand{
		Key: "foo",
		Val: []byte("bar"),
	}

	if cmd.(SetCommand).Key != expected.Key || string(cmd.(SetCommand).Val) != string(expected.Val) {
		t.Errorf("expected %s, got %s\n", expected, cmd)
	}

}
