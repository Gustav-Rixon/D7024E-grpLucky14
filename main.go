package main

import (
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	ID *KademliaID
}

func NewNode() Node {
	Id := NewRandomKademliaID()
	return Node{Id}
}

// Returns random number, used in Kademlia ID generation
func getRandNum() int {
	r := rGen.Intn(256)
	return r
}

// Random number generator, use to get random numbers between nodes
var rGen *rand.Rand

// The node itself
var node Node

func main() {
	// initialize randomization of ID
	randSource := rand.NewSource(time.Now().UnixNano())
	rGen = rand.New(randSource)
	node.ID = NewRandomKademliaID()

	// initialize network settings, communicate via port 80
	initNetwork(80)

	if netInfo.localIPAddr.Mask(net.IPv4Mask(0, 0, 255, 255)).String() == "0.0.0.2" {
		// Lowest IP address, assign supernode
		go listen()
	} else {
		go sendLoop()
	}

	for {
		//fmt.Println("Alive") // Debug printout to ensure node is alive
		time.Sleep(time.Second / 2)
	}
}

func sendLoop() {
	networkPrefix1, _ := strconv.Atoi(strings.Split(netInfo.localIPAddr.String(), ".")[0])
	networkPrefix2, _ := strconv.Atoi(strings.Split(netInfo.localIPAddr.String(), ".")[1])
	supernodeAddr := net.IPv4(byte(networkPrefix1), byte(networkPrefix2), 0, 2)
	for {
		// Forever ping the supernode
		sendPing(supernodeAddr)

		time.Sleep(2 * time.Second)
	}
}
