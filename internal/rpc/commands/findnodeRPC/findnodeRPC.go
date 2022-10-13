package findnodeRPC

import (
	"errors"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"os"
	"strconv"

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

	K, err := strconv.Atoi(os.Getenv("K"))
	if err != nil {
		log.Error().Msgf("Failed to convert env variable ALPHA from string to int: %s", err)
	}

	// Responde with k clossets nodes
	kClosest := node.FindKClosest(kademliaid.FromString(*targetID.targetID), targetID.requestor.ID, K)
	content := contact.SerializeContacts(kClosest)
	node.Network.SendFindContactRespMessage(node.ID, targetID.requestor.Address, targetID.rpcId, &content)

}

func (fn *FindNodeRPC) ParseOptions(options *[]string) error {
	if len(*options) == 0 {
		return errors.New("Recieved empty FIND_NODE RPC, Missing ID argument")
	}
	fn.targetID = &(*options)[0]
	return nil
}
