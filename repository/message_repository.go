// repository/message_repository.go
package repository

import (
	"context"
	"microservice/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type MessageRepository struct {
	DB *pgxpool.Pool
}

func (repo *MessageRepository) SaveMessage(msg *models.Message) error {
	_, err := repo.DB.Exec(context.Background(), "INSERT INTO messages (content, status) VALUES ($1, $2)", msg.Content, msg.Status)
	return err
}

func (repo *MessageRepository) GetProcessedMessagesStats() (map[string]int, error) {
	rows, err := repo.DB.Query(context.Background(), "SELECT status, COUNT(*) FROM messages GROUP BY status")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := make(map[string]int)
	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return nil, err
		}
		stats[status] = count
	}

	return stats, nil
}
