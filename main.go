package main

import (
	. "kademlia/internal/kademliaid"
	"kademlia/internal/network"
	"kademlia/internal/node"
	. "kademlia/internal/node"
	"kademlia/internal/routingtable"
	"net"
	"strconv"
	"strings"
	"time"
)

func getDistance(NodeA []byte, NodeB []byte) KademliaID {
	result := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result[i] = NodeA[i] ^ NodeB[i]
	}
	return result
}

func main() {
	//ROUTINGTABLE TESTING CODE
	node.CreateSelf(NewRandomKademliaID(), network.GetOutboundIP())
	routingtable.NewRoutingTable(*GetNode())

	// initialize network settings, communicate via port 80
	network.InitNetwork(80, 90)

	// If the local ip address does not end with 0.2, it is not the supernode and should enter the sendLoop
	if network.NetInfo.LocalIPAddr.Mask(net.IPv4Mask(0, 0, 255, 255)).String() != "0.0.0.2" {
		go sendLoop()
	}

	go network.Listen()

	for {
		//fmt.Println("Alive") // Debug printout to ensure node is alive
		time.Sleep(time.Second / 10)
	}
}

// Basic test function for constantly pinging the supernode
func sendLoop() {
	networkPrefix1, _ := strconv.Atoi(strings.Split(network.NetInfo.LocalIPAddr.String(), ".")[0])
	networkPrefix2, _ := strconv.Atoi(strings.Split(network.NetInfo.LocalIPAddr.String(), ".")[1])

	// Construct supernode address (xxx.xxx.0.2)
	supernodeAddr := net.IPv4(byte(networkPrefix1), byte(networkPrefix2), 0, 2)
	for {
		// Forever ping the supernode
		network.SendPing(supernodeAddr)
		time.Sleep(time.Second * 3)
	}
}
