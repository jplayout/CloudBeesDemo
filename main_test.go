package main

import (
	"bytes"
	"testing"
)

func TestReadName(t *testing.T) {
	input := bytes.NewBuffer([]byte("Dude"))
	name := readName(input)
	if name != "Dude" {
		t.Error("Expected Dude, got", name)
	}
}

func TestEmptyName(t *testing.T) {
	input := bytes.NewBuffer([]byte(""))
	name := readName(input)
	if name != "You" {
		t.Error("Expected You, got", name)
	}
}
