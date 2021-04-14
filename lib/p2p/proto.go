package p2p

import (
	"encoding/json"
)

// File that defines message structure

// Message struct includes data from message
type Message struct {
	OP string
}

func parseMessage(data []byte) Message {
	var msg Message
	json.Unmarshal(data, &msg)
	return msg
}

func serializeMessage(msg Message) []byte {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil
	}
	data = append(data, '\n')
	return data
}

var IsAliveMsg = Message{
	OP: "isAlive",
}
var AliveRes = Message{
	OP: "isAliveRes",
}
