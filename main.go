package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
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

// Borrwed .)
// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// Returns random number, used in Kademlia ID generation
func getRandNum() int {
	r := rGen.Intn(256)
	return r
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

// Bucket used for testing
var b *Bucket

func main() {
	NodeId1 := []byte{0, 0}
	NodeId2 := []byte{1, 1}
	m := getDistance(NodeId1, NodeId2)
	fmt.Println(m)
}

/*
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
*/

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
