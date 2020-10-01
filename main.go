package main

import (
	"crypto/sha256"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	h := sha256.New()
	_, err := h.Write([]byte("I am the stone"))
	if err != nil {
		spew.Dump(err)
	}

	spew.Printf("%x", h.Sum(nil))
}
