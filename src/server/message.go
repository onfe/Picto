package server

//Message contains everything the server needs to know about the message
type Message struct {
	SenderID int
	Body     []byte
}

func newMessage(Body []byte, SenderID int) Message {
	return Message{
		SenderID: SenderID,
		Body:     Body,
	}
}
