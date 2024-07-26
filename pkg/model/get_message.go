package model

import "time"

type GetMessage struct {
	ID       int   `json:"id"`
	SenderID int   `json:"sender_id"`
	Sender   *User `json:"sender,omitempty"`
	//RecipientID int    `json:"recipient_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}
