package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\nline1\nline2")
	exp := 3
	res := count(b, true)
	if exp != res {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}
