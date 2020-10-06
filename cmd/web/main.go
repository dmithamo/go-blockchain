package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	if err := initializeRouter(); err != nil {
		log.Fatal("Could not initialize server\n", err)
	}
}

// initializeRouter configures a mux server, adds routes and starts the server up
func initializeRouter() error {

	mux := mux.NewRouter()
	mux.HandleFunc("/", retrieveAllRecords).Methods("GET")
	mux.HandleFunc("/", addRecord).Methods("POST")

	err := godotenv.Load()
	if err != nil {
		return err
	}

	port := os.Getenv("PORT")
	server := &http.Server{
		Addr:           port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Listening at http(s)://localhost%s\n", port)
	err = server.ListenAndServe()
	return err
}
