package p2p

import (
	"net"
)

type Node struct {
	IpAddr     *net.TCPAddr
	Connection *net.TCPConn
}

func CreateLocalNode(port int) error {
	//portstr := fmt.Sprintf(":%d", port)
	localAddr := &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: port,
	}
	in, err := net.ListenTCP("tcp4", localAddr)
	if err != nil {
		return err
	}
	go listenForConnections(in)
	return nil
}

func IsAlive(addr *net.TCPAddr) (bool, error) {
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		return false, err
	}
	msg := serializeMessage(IsAliveMsg)
	conn.Write(msg)
	res := readMsg(conn)

	response := parseMessage(res)
	conn.Close()
	return response.OP == IsAliveRes.OP, nil
}

// DiscoverNodes finds initial nodes in the p2p network
func DiscoverNodes() []*Node {
	// Read addressbook
	// TODO
	return nil
}

// ConnectTo starts a p2p connection to a specific node
func (node *Node) ConnectTo() error {
	conn, err := net.DialTCP("tcp4", nil, node.IpAddr)
	if err != nil {
		return err
	}
	node.Connection = conn
	return nil
}
