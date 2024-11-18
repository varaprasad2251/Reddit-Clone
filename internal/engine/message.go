package engine

type Message struct {
	ID        int
	Content   string
	SenderID  int
	ReceiverID int
}

func (e *Engine) SendMessage(senderID, receiverID int, content string) *Message {
	message := &Message{
		ID:         len(e.Messages) + 1, // Use a message list or counter for unique IDs
		Content:    content,
		SenderID:   senderID,
		ReceiverID: receiverID,
	}
	return message
}
