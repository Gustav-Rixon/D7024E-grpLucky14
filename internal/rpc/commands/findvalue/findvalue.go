package findvalue

import (
	"errors"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

type FindValue struct {
	hash          *kademliaid.KademliaID
	rpcId         *kademliaid.KademliaID //dont know???????
	senderAddress *contact.Contact       //Prob
}

func New(sender *contact.Contact, rpcId *kademliaid.KademliaID) FindValue {
	return FindValue{senderAddress: sender, rpcId: rpcId}
}

func (find FindValue) Execute(node *node.Node) {
	log.Trace().Msg("Executing Find_VALUE RPC")

	if value := node.DataStore.Get(*find.hash); value != "" {
		log.Debug().Str("Value", value).Str("Hash", find.hash.String()).Msg("Key Found")
		response := "Value=" + value
		node.Network.SendFindDataRespMessage(node.ID, find.senderAddress.Address, find.rpcId, &response)
	} else {
		panic("ajajajajaja")
	}
}

func (find FindValue) ParseOptions(options *[]string) error {
	if (len(*options)) == 0 {
		return errors.New("No hash")
	}

	return nil
}
