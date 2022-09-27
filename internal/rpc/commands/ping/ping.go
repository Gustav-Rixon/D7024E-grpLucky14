package ping

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

// TODO RENAME ALL RPC COMMANDS TO <NAME>RPC
type Ping struct {
	senderID      *kademliaid.KademliaID
	senderAddress *address.Address
	rpcId         *kademliaid.KademliaID
}

func New(senderID *kademliaid.KademliaID, senderAddress *address.Address, rpcId *kademliaid.KademliaID) Ping {
	return Ping{senderID: senderID, senderAddress: senderAddress, rpcId: rpcId}
}

func (ping Ping) Execute(node *node.Node) {
	log.Trace().Msg("Executing PING RPC")
	// Update routing table
	// Respond with pong and add it to its table
	node.Network.SendPongMessage(node.ID, ping.senderAddress, ping.rpcId)
}

func (ping Ping) ParseOptions(options *[]string) error {
	// Ping takes no options
	return nil
}
