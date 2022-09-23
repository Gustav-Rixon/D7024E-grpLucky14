package node

import (
	"kademlia/internal/kademliaid"
	"net"
)

// Used in main to call on NewRandomKademliaID function
type Node struct {
	ID       [kademliaid.IDLength]byte
	IP       net.IP
	distance []byte
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

func (node Node) GetDistance() []byte {
	return node.distance
}

func NewNode(id [kademliaid.IDLength]byte, ip net.IP) Node {
	Id := kademliaid.NewKademliaID(id)
	//fmt.Println("Successfully created instance of Kademlia ID: ", *Id, " With IP: ", ip.String())
	tempdist := []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	return Node{Id, ip, tempdist}
}

// CalcDistance returns a new instance of a KademliaID that is built
// through a bitwise XOR operation betweeen kademliaID and target
func (n Node) CalcDistance(target [kademliaid.IDLength]byte) [kademliaid.IDLength]byte {
	result := kademliaid.KademliaID{}
	for i := 0; i < kademliaid.IDLength; i++ {
		result[i] = n.ID[i] ^ target[i]
		n.distance[i] = result[i]
	}

	return result
}
