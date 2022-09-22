package main

import (
	"fmt"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	. "kademlia/internal/node"
	"net"
)

func main() {
	CreateSelf(kademliaid.NewRandomKademliaID(), net.IPv4(0, 0, 0, 0)) // init table
	fmt.Println(*node.GetNode())
}
