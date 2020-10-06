package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// BPM - beats per minute - is the type of the 'actual' data stored on a Block
type BPM int

type RecordParams struct {
	BPM int `json:"bpm,omitempty"`
}

// Record defines a data record stored on the blockchain
type Record struct {
	Index     int       // index of this record on the chain
	Timestamp time.Time // timestamp when this record was created
	BPM       BPM       // data stored on this record (Beats Per Minute)
	Hash      string    // SHA256 id of this record, formed from hashing its data
	PrevHash  string    // SHA256 id of prev record, formed from hashing prev record's data
}

// computeOwnHash generates the unique hash for a record
func (r *Record) computeOwnHash() error {
	record := fmt.Sprintf("%d::%d::%d::%s", r.Index, r.Timestamp.Unix(), r.BPM, r.PrevHash)

	// initialize hasher
	h := sha256.New()

	_, err := h.Write([]byte(record))
	if err != nil {
		return err
	}

	r.Hash = string(h.Sum(nil))
	return nil
}
