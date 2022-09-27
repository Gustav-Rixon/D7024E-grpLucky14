package findnode

import (
	"errors"
	"kademlia/internal/address"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
)

// IN NSLLOCKUP SENDERID = TARGETID NOTE FOR GURX
type FindNodeRPC struct {
	senderID      *kademliaid.KademliaID
	senderAddress *address.Address
	targetID      *kademliaid.KademliaID
	rpcId         *kademliaid.KademliaID
}

func New(senderID *kademliaid.KademliaID, senderAddress *address.Address, targetID *kademliaid.KademliaID, rpcId *kademliaid.KademliaID) FindNodeRPC {
	return FindNodeRPC{senderID: senderID, senderAddress: senderAddress, targetID: targetID, rpcId: rpcId}
}

func (targetID *FindNodeRPC) Execute(node *node.Node) ([]contact.Contact, error) {
	candidats := node.FindKClosest(targetID.targetID, node.ID, 3)
	return candidats, nil
}

func (targetID *FindNodeRPC) ParseOptions(options []string) error {
	if len(options) < 1 {
		return errors.New("Missing hash")
	}
	targetID.targetID = kademliaid.FromString(options[0])
	return nil
}

func (hash *FindNodeRPC) PrintUsage() string {
	return "USAGE: get <hash>"
}
