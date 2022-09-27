package findnode

import (
	"errors"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
)

// IN NSLLOCKUP SENDERID = TARGETID NOTE FOR GURX
type FindNodeRPC struct {
	targetID  *string          //kommer bara vara en string som input ellr?
	requestor *contact.Contact // Contact infon om vem som fråga
	rpcId     *kademliaid.KademliaID
}

func New(requestor *contact.Contact, rpcId *kademliaid.KademliaID) *FindNodeRPC {
	return &FindNodeRPC{requestor: requestor, rpcId: rpcId}
}

func (targetID *FindNodeRPC) Execute(node *node.Node) {
	candidats := node.FindKClosest(kademliaid.FromString(*targetID.targetID), targetID.rpcId, 3)
	contact := contact.SerializeContacts(candidats)
	node.Network.SendFindDataRespMessage(node.ID, targetID.requestor.Address, targetID.rpcId, &contact)
}

func (targetID *FindNodeRPC) ParseOptions(options *[]string) error {
	if len(*options) < 1 {
		return errors.New("Missing hash")
	}
	return nil
}

func (hash *FindNodeRPC) PrintUsage() string {
	return "USAGE: get <hash>"
}
