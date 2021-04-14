package main

import (
	"capycoin/lib/database"
	"capycoin/lib/p2p"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World")
	database.OpenConnection()

	contact := &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 9999,
	}
	database.AddContact(contact)
	err := p2p.CreateLocalNode(6969)
	if err != nil {
		panic(err)
	}

	contacts := database.QueryAddressbook()
	fmt.Println(len(contacts))
	database.CloseConn()
}
