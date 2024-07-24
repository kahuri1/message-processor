package model

type Message struct {
	SenderID int64 `json:"sender_id"`
	//RecipientID int64  `json:"recipient_id"`
	Content string `json:"content"`
}
