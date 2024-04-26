package main

import (
	"testing"
)

func TestParseCommad(t *testing.T) {
	msg := "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"

	cmd, err := parseCommand(msg)
	if err != nil {
		t.Fatal(err)
	}

	expected := SetCommand{
		key: "foo",
		val: "bar",
	}

	if cmd != expected {
		t.Errorf("expected %s, got %s\n", expected, cmd)
	}

}
