package blockchain

import "time"

// Chain is the slice of all members of the blockchain
var Chain []*Record

// GenerateNewRecord initializes a record and appends it to the Chain
func GenerateNewRecord(bpm BPM) (*Record, error) {
	newRecord := &Record{
		Timestamp: time.Now(),
		BPM:       bpm,
	}

	prevRecord := getLastestRecord()
	if prevRecord == nil {
		newRecord.Index = 0
		newRecord.PrevHash = ""
	} else {
		newRecord.Index = prevRecord.Index + 1
		newRecord.PrevHash = prevRecord.Hash
	}

	err := newRecord.computeOwnHash()
	if err != nil {
		return nil, err
	}

	Chain = append(Chain, newRecord)
	return newRecord, nil
}

// IsRecordValid checks that a record is an authentic member of the blockchain
func IsRecordValid(r *Record) bool {
	if r.Index == 0 {
		return r.PrevHash == ""
	}

	// check prevHash is indeed the hash of the index-1 member of the chain
	// also check that index of record is 1+(index of prev record)
	return r.PrevHash == Chain[r.Index-1].Hash && r.Index == Chain[r.Index-1].Index+1
}

// getLastestRecord convenience fn for retrieving prev record
func getLastestRecord() *Record {
	if len(Chain) == 0 {
		return nil
	}

	return Chain[len(Chain)-1]
}
