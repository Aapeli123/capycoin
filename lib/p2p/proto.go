package p2p

// File that defines message structure

// Message struct includes data from message
type Message struct {
	OPCode string
}

func parseMessage(data []byte) Message {
	return Message{}
}
