package server

import "encoding/json"

//Message contains everything the server needs to know about the message
type Message struct {
	SenderID int    `json:"SenderID"`
	Body     []byte `json:"Body"`
}

func newMessage(Body []byte, SenderID int) Message {
	return Message{
		SenderID: SenderID,
		Body:     Body,
	}
}

func (m *Message) getEventData() []byte {
	data, _ := json.Marshal(MessageEvent{
		Event:     Event{Event: "message"},
		UserIndex: m.SenderID,
		Message:   m.Body,
	})
	return data
}
