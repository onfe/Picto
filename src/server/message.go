package server

//Message contains everything the server needs to know about the message
type Message struct {
	senderID int
	body     []byte
}

func newMessage(message []byte, senderID int) Message {
	return Message{
		senderID: senderID,
		body:     message,
	}
}
