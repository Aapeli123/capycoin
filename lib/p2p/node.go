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
	in.Close()
	return nil
}

// DiscoverNodes finds initial nodes in the p2p network
func DiscoverNodes() []*Node {
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
