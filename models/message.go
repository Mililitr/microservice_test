// models/message.go
package models

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

type MessageRepository struct {
	DB *pgxpool.Pool
}

func (repo *MessageRepository) SaveMessage(msg *Message) error {
	_, err := repo.DB.Exec(context.Background(), "INSERT INTO messages (content, status) VALUES ($1, $2)", msg.Content, msg.Status)
	return err
}
