package main

import (
	"capycoin/lib/p2p"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	err := p2p.CreateLocalNode(6969)
	if err != nil {
		panic(err)
	}
}
