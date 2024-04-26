package protocol

import (
	"testing"
)

func TestParseCommad(t *testing.T) {
	msg := "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"

	cmd, err := ParseCommand(msg)
	if err != nil {
		t.Fatal(err)
	}

	expected := SetCommand{
		Key: "foo",
		Val: "bar",
	}

	if cmd != expected {
		t.Errorf("expected %s, got %s\n", expected, cmd)
	}

}
