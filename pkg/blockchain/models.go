package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// BPM - beats per minute - is the type of the 'actual' data stored on a Block
type BPM int

// Record defines a data record stored on the blockchain
type Record struct {
	Index     int       // index of this record on the chain
	Timestamp time.Time // timestamp when this record was created
	BPM       BPM       // data stored on this record (Beats Per Minute)
	Hash      string    // SHA256 id of this record, formed from hashing its data
	PrevHash  string    // SHA256 id of prev record, formed from hashing prev record's data
}

// computeOwnHash generates the unique hash for a record
func (b *Record) computeOwnHash() error {
	record := fmt.Sprintf("%d%d%d%s", b.Index, b.Timestamp.Unix(), b.BPM, b.PrevHash)

	// initialize hasher
	h := sha256.New()

	_, err := h.Write([]byte(record))
	if err != nil {
		return err
	}

	b.Hash = string(h.Sum(nil))
	return nil
}
