package main

import (
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

func NewNode(id [IDLength]byte, ip net.IP) Node {
	Id := NewKademliaID(id)
	//fmt.Println("Successfully created instance of Kademlia ID: ", *Id, " With IP: ", ip.String())
	return Node{Id, ip}
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

// Bucket used for testing
var b *Bucket

func main() {
	// initialize randomization of ID
	randSource := rand.NewSource(time.Now().UnixNano())
	rGen = rand.New(randSource)
	node.ID = NewRandomKademliaID()
	node.IP = getOutboundIP()

	//BUCKET TESTING CODE
	b = newBucket()

	// initialize network settings, communicate via port 80
	initNetwork(80, 90)

	if netInfo.localIPAddr.Mask(net.IPv4Mask(0, 0, 255, 255)).String() != "0.0.0.2" {
		go sendLoop()
	}

	go listen()

	for {
		//fmt.Println("Alive") // Debug printout to ensure node is alive
		time.Sleep(time.Second / 10)
	}
}

func sendLoop() {
	networkPrefix1, _ := strconv.Atoi(strings.Split(netInfo.localIPAddr.String(), ".")[0])
	networkPrefix2, _ := strconv.Atoi(strings.Split(netInfo.localIPAddr.String(), ".")[1])

	// Find supernode address (xxx.xxx.0.2)
	supernodeAddr := net.IPv4(byte(networkPrefix1), byte(networkPrefix2), 0, 2)
	for {
		// Forever ping the supernode
		sendPing(supernodeAddr)
		time.Sleep(time.Second * 3)
	}
}
