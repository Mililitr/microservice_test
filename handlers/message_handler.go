// handlers/message_handler.go
package handlers

import (
	"encoding/json"
	"net/http"

	"microservice/kafka"
	"microservice/models"
	"microservice/repository"

	"github.com/IBM/sarama"
)

type MessageHandler struct {
	Repo     *repository.MessageRepository
	Producer sarama.SyncProducer
}

func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Repo.SaveMessage(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := kafka.SendMessage(h.Producer, "your_topic", msg.Content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MessageHandler) GetProcessedMessagesStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.Repo.GetProcessedMessagesStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stats)
}
