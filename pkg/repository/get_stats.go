package repository

import "fmt"

func (r *Repository) GetStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Получение общего количества сообщений
	var totalMessages int
	err := r.db.QueryRow("SELECT COUNT(*) FROM messages").Scan(&totalMessages)
	if err != nil {
		return nil, fmt.Errorf("failed to get total messages: %w", err)
	}
	stats["total_messages"] = totalMessages

	// Получение количества сообщений по отправителям
	rows, err := r.db.Query("SELECT sender_id, COUNT(*) FROM messages GROUP BY sender_id")
	if err != nil {
		return nil, fmt.Errorf("failed to get messages by sender: %w", err)
	}
	defer rows.Close()

	senderStats := make(map[int]int) // sender_id -> message_count
	for rows.Next() {
		var senderID int
		var messageCount int
		if err := rows.Scan(&senderID, &messageCount); err != nil {
			return nil, err
		}
		senderStats[senderID] = messageCount
	}
	stats["messages_by_sender"] = senderStats

	return stats, nil
}
