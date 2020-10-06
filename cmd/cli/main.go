package main

import (
	"log"

	chain "github.com/dmithamo/go-blockchain/pkg/blockchain"
)

func main() {

	bpms := []int{100, 20}
	for i := 0; i < len(bpms); i++ {
		_, err := chain.GenerateNewRecord(chain.BPM(bpms[i]))
		checkErrorHelper(err)
	}

}

func checkErrorHelper(err error) {
	if err != nil {
		log.Fatal("could new initialize record:", err)
	}
}
