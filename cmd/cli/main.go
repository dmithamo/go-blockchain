package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	chain "github.com/dmithamo/go-blockchain/pkg/blockchain"
)

func main() {

	bpms := []int{100, 82, 90, 99, 85}
	for i := 0; i < len(bpms); i++ {
		_, err := chain.GenerateNewRecord(chain.BPM(bpms[i]))
		checkErrorHelper(err)
	}

	spew.Dump(chain.Chain[:])
}

func checkErrorHelper(err error) {
	if err != nil {
		log.Fatal("could new initialize record:", err)
	}
}
