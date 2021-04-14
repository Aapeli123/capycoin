package main

import (
	"capycoin/lib/database"
	"capycoin/lib/p2p"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	database.OpenConnection()
	/*
		contact := &net.TCPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 9999,
		}

		database.AddContact(contact)
	*/
	//	err := p2p.CreateLocalNode(6969)

	contacts := database.QueryAddressbook()
	fmt.Println(p2p.IsAlive(contacts[0]))
	//	if err != nil {
	//panic(err)
	//}
	//for {

	//	}
	fmt.Println(len(contacts))
	database.CloseConn()
}
