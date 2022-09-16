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
	ACK
	return_nodes // Created when node receives a find_node or find_value Packet
	return_value // Created when node receives a find_value Packet and it has the stored value
)

// Catch-all struct for all kinds of packets
// NOTE: Some fields are only used in certain PacketTypes
type Packet struct {
	pType PacketType

	ID       [20]byte // The ID of the original sender
	IP       net.IP   // The IP of the original sender
	targetID [20]byte // The ID that the original sender wants the packet to reach

	keyVal []byte // Holds the 20-byte key in find_value Packets, but is also used to hold the value in store and return_value Packets

	nodes []Node // Array used for storing node info when responding to a find_node packet
}

func createPingPacket(thisNode Node) Packet {
	var p Packet
	p.pType = ping

	p.ID = thisNode.ID
	p.IP = thisNode.IP

	return p
}

func createFindNodePacket(thisNode Node, targetID [20]byte) Packet {
	var p Packet
	p.pType = find_node

	p.ID = thisNode.ID
	p.IP = thisNode.IP
	p.targetID = targetID

	return p
}

func createFindValuePacket(thisNode Node, targetID [20]byte, key [20]byte) Packet {
	var p Packet
	p.pType = find_value

	p.ID = thisNode.ID
	p.IP = thisNode.IP
	p.targetID = targetID
	p.keyVal = key[:]

	return p
}

func createStorePacket(thisNode Node, value []byte) Packet {
	var p Packet
	p.pType = store

	p.ID = thisNode.ID
	p.IP = thisNode.IP
	p.keyVal = value

	return p
}

func createACKPacket(thisNode Node) Packet {
	var p Packet
	p.pType = ACK

	p.ID = thisNode.ID
	p.IP = thisNode.IP

	return p
}

func createReturnNodesPacket(thisNode Node, nodes []Node) Packet {
	var p Packet
	p.pType = return_nodes

	p.ID = thisNode.ID
	p.IP = thisNode.IP
	p.nodes = nodes

	return p
}

func createReturnValuePacket(thisNode Node, value []byte) Packet {
	var p Packet
	p.pType = return_nodes

	p.ID = thisNode.ID
	p.IP = thisNode.IP
	p.keyVal = value

	return p
}
