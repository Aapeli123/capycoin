package p2p

// Store all connections in AddressBook
// When connecting to a new node share address book with them and add new nodes to addressbook
var AddressBook []*Node

// GetAddressData Converts addresbook to array of addressdata which is more easily sent over network
func GetAddressData() []AddressData {
	data := []AddressData{}

	for _, addr := range AddressBook {
		a := AddressData{
			IP:   addr.IpAddr.String(),
			Port: addr.IpAddr.Port,
		}
		data = append(data, a)
	}
	return data
}

func Broadcast(msg Message) {

}
