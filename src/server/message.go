package server

//Message contains everything the server needs to know about the message
type Message struct {
	SenderID string `json:"SenderID"`
	Body     []byte `json:"Body"`
}

func newMessage(Body []byte, SenderID string) Message {
	return Message{
		SenderID: SenderID,
		Body:     Body,
	}
}
