package kademliaid

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// the static number of bytes in a KademliaID
const IDLength int = 20

// type definition of a KademliaID
type KademliaID [IDLength]byte

// Förmodligen för testing?
// NewKademliaID returns a new instance of a KademliaID based on the string input
func NewKademliaID(id [IDLength]byte) [IDLength]byte {
	newKademliaID := KademliaID{}
	for i := 0; i < IDLength; i++ {
		newKademliaID[i] = id[i]
	}

	return newKademliaID
}

// Random number generator, use to get random numbers between nodes
var rGen *rand.Rand

// Returns random number, used in Kademlia ID generation
func getRandNum() int {
	r := rGen.Intn(256)
	return r
}

// NewRandomKademliaID returns a new instance of a random KademliaID,
// change this to a better version if you like
func NewRandomKademliaID() [20]byte {
	// First time init
	if rGen == nil {
		randSource := rand.NewSource(time.Now().UnixNano())
		rGen = rand.New(randSource)
	}

	newKademliaID := KademliaID{}
	for i := 0; i < IDLength; i++ {
		newKademliaID[i] = uint8(getRandNum())
	}
	return newKademliaID
}

// Less returns true if kademliaID < otherKademliaID (bitwise)
func (kademliaID KademliaID) Less(otherKademliaID *KademliaID) bool {
	for i := 0; i < IDLength; i++ {
		if kademliaID[i] != otherKademliaID[i] {
			return kademliaID[i] < otherKademliaID[i]
		}
	}
	return false
}

// Equals returns true if kademliaID == otherKademliaID (bitwise)
func (kademliaID KademliaID) Equals(otherKademliaID *KademliaID) bool {
	for i := 0; i < IDLength; i++ {
		if kademliaID[i] != otherKademliaID[i] {
			return false
		}
	}
	return true
}

// String returns a simple string representation of a KademliaID
func (kademliaID *KademliaID) String() string {
	return hex.EncodeToString(kademliaID[0:IDLength])
}
