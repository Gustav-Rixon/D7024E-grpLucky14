package pong

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
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
	log.Trace().Msg("Executing PONG RPC")
}

func (pong Pong) ParseOptions(options *[]string) error {
	// Pong takes no options
	return nil
}
