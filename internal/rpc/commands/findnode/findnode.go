package findnode

import (
	"errors"
	"fmt"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
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

// [{98521a170970a249b276ee78eecc60853ff1c9c1 172.18.0.4:8888 0xc0000c8678} {9514e18b679622b8d59991a6298559cb03099d64 172.18.0.2:8888 0xc0000c8660}]
func (targetID *FindNodeRPC) Execute(node *node.Node) {
	log.Trace().Msg("Executing FIND_NODE RPC")
	candidats := node.FindKClosest(kademliaid.FromString(*targetID.targetID), targetID.rpcId, 3)
	contact := contact.SerializeContacts(candidats)
	fmt.Println(contact)
	node.Network.SendFindContactRespMessage(node.ID, targetID.requestor.Address, targetID.rpcId, &contact)
}

func (targetID *FindNodeRPC) ParseOptions(options *[]string) error {
	if len(*options) < 1 {
		return errors.New("Missing hash")
	}
	targetID.targetID = &(*options)[0]
	return nil
}

func (hash *FindNodeRPC) PrintUsage() string {
	return "USAGE: get <hash>"
}
