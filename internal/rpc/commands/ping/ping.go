package ping

import (
	"kademlia/internal/address"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

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
	node.RoutingTable.AddContact(contact.NewContact(ping.senderID, ping.senderAddress))
	// Respond with pong
	node.Network.SendPongMessage(node.ID, ping.senderAddress, ping.rpcId)
}

func (ping Ping) ParseOptions(options *[]string) error {
	// Ping takes no options
	return nil
}
