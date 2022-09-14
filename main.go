package main

import (
	. "kademlia/internal/bucket"
	. "kademlia/internal/kademliaid"
	. "kademlia/internal/node"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

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

func getDistance(NodeA []byte, NodeB []byte) KademliaID {
	result := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result[i] = NodeA[i] ^ NodeB[i]
	}
	return result
}

// The node itself
var node Node

// Bucket used for testing
var b *Bucket

func main() {
	// initialize randomization of ID
	//randSource := rand.NewSource(time.Now().UnixNano())
	//rGen = rand.New(randSource)
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
