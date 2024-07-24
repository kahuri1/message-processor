package todo

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}

type Message struct {
	ID          int    `json:"id"`
	SenderID    int    `json:"sender_id"`
	Sender      *User  `json:"sender,omitempty"`
	RecipientID int    `json:"recipient_id"`
	Content     string `json:"content"`
	CreatedAt   int64  `json:"created_at"`
	Status      string `json:"status"`
}
