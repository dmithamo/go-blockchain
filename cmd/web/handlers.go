package main

import (
	"encoding/json"
	"net/http"

	"github.com/dmithamo/go-blockchain/pkg/blockchain"
)

// retrieveAllRecords reads the chain and returns all records on it
func retrieveAllRecords(w http.ResponseWriter, r *http.Request) {
	// read the chain
	records := blockchain.Chain
	if len(records) == 0 {
		writeJSONResponseHelper(w, "No records found", http.StatusNotFound)
		return
	}

	writeJSONResponseHelper(w, records, http.StatusOK)
}

// addRecord inserts a new record into the chain
func addRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var recordParams blockchain.RecordParams
	decodeErr := json.NewDecoder(r.Body).Decode(&recordParams)
	if decodeErr != nil {
		handleInternalErrorHelper(decodeErr, w)
		return
	}

	record, err := blockchain.GenerateNewRecord(blockchain.BPM(recordParams.BPM))
	if err != nil {
		handleInternalErrorHelper(err, w)
		return
	}

	writeJSONResponseHelper(w, record, http.StatusCreated)
}

// writeJSONResponseHelper prints out a json response
func writeJSONResponseHelper(w http.ResponseWriter, response interface{}, status int) {
	resAsJSON, err := json.MarshalIndent(response, "", "	")
	if err != nil {
		handleInternalErrorHelper(err, w)
		return
	}

	w.WriteHeader(status)
	_, err = w.Write(resAsJSON)

	if err != nil {
		handleInternalErrorHelper(err, w)
	}
}

// handleInternalErrorHelper handles internal errs gracefully
func handleInternalErrorHelper(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusInternalServerError)

}
