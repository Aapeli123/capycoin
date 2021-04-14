package p2p

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn *net.TCPConn) {
	// var buf bytes.Buffer
	reader := bufio.NewReader(conn)
	data, _ := reader.ReadBytes('\n')
	/*io.Copy(&buf, conn)
	data := buf.Bytes()
	*/
	msg := parseMessage(data)
	fmt.Println(msg)
	switch msg.OP {
	case "isAlive":
		res := serializeMessage(AliveRes)
		conn.Write(res)
	default:
		fmt.Println("Unknown operation: " + msg.OP)
	}
}

func listenForConnections(in *net.TCPListener) {
	for {
		conn, err := in.AcceptTCP()
		if err != nil {
			fmt.Println("Error accepting connection: " + err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func readMsg(conn *net.TCPConn) []byte {
	reader := bufio.NewReader(conn)
	data, _ := reader.ReadBytes('\n')
	return data
}
