package node

import (
	"kademlia/internal/kademliaid"
	"net"
)

// Used in main to call on NewRandomKademliaID function
type Node struct {
	ID [kademliaid.IDLength]byte
	IP net.IP
}

// This container's node
var node Node

// Initiates the node
func CreateSelf(id [kademliaid.IDLength]byte, ip net.IP) {
	// Only initialize if ID is uninitialized
	if node.ID == [20]byte{} {
		node = NewNode(id, ip)
	} else {
		panic("WHY WOULD YOU TRY TO INIT TWICE")
	}
}

func GetNode() *Node {
	return &node
}

func NewNode(id [kademliaid.IDLength]byte, ip net.IP) Node {
	Id := kademliaid.NewKademliaID(id)
	//fmt.Println("Successfully created instance of Kademlia ID: ", *Id, " With IP: ", ip.String())
	return Node{Id, ip}
}

// CalcDistance returns a new instance of a KademliaID that is built
// through a bitwise XOR operation betweeen kademliaID and target
func (node Node) CalcDistance(target [kademliaid.IDLength]byte) [kademliaid.IDLength]byte {
	result := kademliaid.KademliaID{}
	for i := 0; i < kademliaid.IDLength; i++ {
		result[i] = node.ID[i] ^ target[i]
	}
	return result
}
