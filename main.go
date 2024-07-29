// main.go
package main

import (
	"log"
	"net/http"

	"microservice/config"
	"microservice/handlers"
	"microservice/kafka"
	"microservice/repository"

	"github.com/gorilla/mux"
)

func main() {
	db, err := config.GetDBConnection()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	repo := &repository.MessageRepository{DB: db}

	brokers := []string{"localhost:9092"}
	producer, err := kafka.NewProducer(brokers)
	if err != nil {
		log.Fatalf("Could not create Kafka producer: %v", err)
	}
	defer producer.Close()

	handler := &handlers.MessageHandler{Repo: repo, Producer: producer}

	r := mux.NewRouter()
	r.HandleFunc("/messages", handler.CreateMessage).Methods("POST")
	r.HandleFunc("/messages/stats", handler.GetProcessedMessagesStats).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
