package pong

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
)

type Pong struct {
	senderID      *kademliaid.KademliaID
	senderAddress *address.Address
	rpcId         *kademliaid.KademliaID
}

func New(senderID *kademliaid.KademliaID, senderAddress *address.Address, rpcId *kademliaid.KademliaID) Pong {
	return Pong{senderID: senderID, senderAddress: senderAddress, rpcId: rpcId}
}

func (pong Pong) Execute(node *node.Node) {
	// MUST BE EMPTY OR CHOASSS
}

func (pong Pong) ParseOptions(options *[]string) error {
	// Pong takes no options
	return nil
}
