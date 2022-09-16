package main

import "net"

// Go version of an enum:
type PacketType uint8

const (
	// Requesting packets:
	ping PacketType = iota
	find_node
	find_value
	store

	// Returning packets:
	return_nodes // Created when node receives a find_node or find_value Packet
	return_value // Created when node receives a find_value Packet and it has the stored value
)

type Packet struct {
	pType    PacketType
	ID       [20]byte  // The ID of the original sender
	IP       net.IP    // The IP of the original sender
	targetID [20]byte  // The ID that the original sender wants the packet to reach
	keyVal   [256]byte // The value (or key) that the original sender wants to store (or read) --- note: (only used in find_value and store packets)

	nodes []Node
}
