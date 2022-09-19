package main

import (
	. "kademlia/internal/kademliaid"
	"kademlia/internal/network"
	. "kademlia/internal/network"
	. "kademlia/internal/node"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

// Used in main to call on NewRandomKademliaID function
type Node struct {
	ID [IDLength]byte
	IP net.IP
}

// Returns random number, used in Kademlia ID generation
func getRandNum() int {
	r := rGen.Intn(256)
	return r
}

// Creates a node instance of itself
func createSelf() Node {
	var me = NewNode(NewRandomKademliaID(), network.GetOutboundIP())
	return me
}

func getDistance(NodeA []byte, NodeB []byte) KademliaID {
	result := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result[i] = NodeA[i] ^ NodeB[i]
	}
	return result
}

// Random number generator, use to get random numbers between nodes
var rGen *rand.Rand

// The node itself
var node Node

// Routing table used for testing
var rt *RoutingTable

func main() {
	//initialize randomization of ID
	randSource := rand.NewSource(time.Now().UnixNano())
	rGen = rand.New(randSource)

	//ROUTINGTABLE TESTING CODE
	node = createSelf()
	rt = NewRoutingTable(node)

	// initialize network settings, communicate via port 80
	initNetwork(80, 90)

	// If the local ip address does not end with 0.2, it is not the supernode and should enter the sendLoop
	if netInfo.localIPAddr.Mask(net.IPv4Mask(0, 0, 255, 255)).String() != "0.0.0.2" {
		go sendLoop()
	}

	go listen()

	for {
		//fmt.Println("Alive") // Debug printout to ensure node is alive
		time.Sleep(time.Second / 10)
	}
}

// Basic test function for constantly pinging the supernode
func sendLoop() {
	networkPrefix1, _ := strconv.Atoi(strings.Split(netInfo.localIPAddr.String(), ".")[0])
	networkPrefix2, _ := strconv.Atoi(strings.Split(netInfo.localIPAddr.String(), ".")[1])

	// Construct supernode address (xxx.xxx.0.2)
	supernodeAddr := net.IPv4(byte(networkPrefix1), byte(networkPrefix2), 0, 2)
	for {
		// Forever ping the supernode
		sendPing(supernodeAddr)
		time.Sleep(time.Second * 3)
	}
}
