package node

import (
	. "kademlia/internal/kademliaid"
	"net"
)

// Used in main to call on NewRandomKademliaID function
type Node struct {
	ID [IDLength]byte
	IP net.UDPAddr
}

func NewNode(id [IDLength]byte, ip net.UDPAddr) Node {
	Id := NewKademliaID(id)
	//fmt.Println("Successfully created instance of Kademlia ID: ", *Id, " With IP: ", ip.String())
	return Node{Id, ip}
}
