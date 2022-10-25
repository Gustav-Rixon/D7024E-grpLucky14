package findvalueRPC

import (
	"errors"
	"fmt"
	"kademlia/internal/constants"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

type FindValueRPC struct {
	hash          *string
	rpcId         *kademliaid.KademliaID //dont know???????
	senderAddress *contact.Contact       //Prob
}

func New(sender *contact.Contact, rpcId *kademliaid.KademliaID) *FindValueRPC {
	return &FindValueRPC{senderAddress: sender, rpcId: rpcId}
}

func (find *FindValueRPC) Execute(node *node.Node) {
	log.Trace().Msg("Executing Find_VALUE RPC")

	key := kademliaid.FromString(*find.hash)

	if value := node.DataStore.GetValue(*key); value != "" {

		fmt.Println("@@@@@@@@@Key Found@@@@@@@@@@@")
		fmt.Println(fmt.Sprintf("Found at node: %s", node.ID))
		fmt.Println(fmt.Sprintf("Item %s", value))
		fmt.Println("@@@@@@@@@Key Found@@@@@@@@@@@")

		log.Debug().Str("Value", value).Str("Hash", *find.hash).Msg("Key Found")
		response := "VALUE=" + value
		node.Network.SendFindDataRespMessage(node.ID, find.senderAddress.Address, find.rpcId, &response)
	} else {
		log.Debug().Str("Hash", *find.hash).Msg("Did not find key")
		closest := node.FindKClosest(key, find.senderAddress.ID, constants.K)
		data := contact.SerializeContacts(closest)
		node.Network.SendFindDataRespMessage(node.ID, find.senderAddress.Address, find.rpcId, &data)
	}
}

func (find *FindValueRPC) ParseOptions(options *[]string) error {
	if (len(*options)) == 0 {
		return errors.New("No hash")
	}

	find.hash = &(*options)[0]
	return nil
}
