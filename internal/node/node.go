package node

import (
	"crypto/sha256"
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

// Hash table used as storage
var storage map[string]string

// Initiates the node if it hasn't been initialized already
func CreateSelf(id [kademliaid.IDLength]byte, ip net.IP) {
	if node.ID == [20]byte{} {
		node = NewNode(id, ip)
		storage = make(map[string]string)
	} else {
		panic("Node already initialized")
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

// Creates and returns a 160-bit hash key for a given string
func GetKey(value string) string {
	key := sha256.Sum256([]byte(value))
	return string(key[:20])
}

// Takes a string and inserts it into this node's hash table
func StoreValue(value string) {
	key := GetKey(value)
	storage[key] = value
}

// Returns the value from the hash table that the key is mapped to if an entry exists, otherwise returns an empty string
func GetValue(key string) (exists bool, value string) {
	value, exists = storage[key]
	if exists {
		return exists, value
	} else {
		// Value not found, does not exist in storage
		return exists, ""
	}
}
