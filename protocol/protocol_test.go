package protocol

import (
	"bytes"
	"testing"
)

func TestParseCommad(t *testing.T) {
	msg := "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"

	cmd, err := parseCommand(bytes.NewBufferString(msg))
	if err != nil {
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
