package node

import (
	"fmt"
	"kademlia/datastore"
	"kademlia/internal/kademliaid"
	"net"
)

// Used in main to call on NewRandomKademliaID function
type Node struct {
	ID        [kademliaid.IDLength]byte
	IP        net.IP
	DataStore datastore.DataStore
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
	fmt.Println("Successfully created instance of Kademlia ID: ", id, " With IP: ", ip.String())
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

func (node *Node) StoreData(value *string) {

}

/*
// Creates and returns a 160-bit hash key for a given string
func GetKey(value string) string {
	key := sha1.Sum([]byte(value))
	return string(key[:20])
}
*/

/*
// Takes a string and inserts it into this node's hash table
func StoreValue(value string) {
	s := value
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	storage[sha1_hash] = value
}
*/

/*
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
*/
