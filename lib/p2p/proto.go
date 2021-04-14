package p2p

import (
	"capycoin/lib/blockchain"
	"encoding/json"
)

// File that defines message structure

// Message struct includes data from message
type Message struct {
	OP            string      // Operation string
	BroadcastType int         // Should data be shared through the network or not
	Data          interface{} // Data, for example Block in the blockchain.
}

// Use if the message should not be spread through the p2p network
const BroadcastKeep = 0

// Use if the message should be spread through the p2p network
const BroadcastSpread = 1

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
	OP:            "isAlive",
	BroadcastType: BroadcastKeep,
}
var IsAliveRes = Message{
	OP:            "isAliveRes",
	BroadcastType: BroadcastKeep,
}

func SendBlockMsg(block *blockchain.Block) Message {
	return Message{
		OP:            "block",
		BroadcastType: BroadcastSpread,
		Data:          block,
	}
}

type AddressData struct {
	IP   string
	Port int
}

func AddrMsg(data []AddressData) Message {
	return Message{
		OP:            "addr",
		BroadcastType: BroadcastKeep,
		Data:          data,
	}
}
