package findvalueRPC

import (
	"errors"
	"fmt"
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

	fmt.Println("###")
	fmt.Println(find.hash)
	fmt.Println("###")

	//if value := node.DataStore.Get(*find.hash); value != "" {
	//	log.Debug().Str("Value", value).Str("Hash", find.hash.String()).Msg("Key Found")
	//	response := "VALUE=" + value
	//	node.Network.SendFindDataRespMessage(node.ID, find.senderAddress.Address, find.rpcId, &response)
	//} else {
	//	k, err := strconv.Atoi(os.Getenv("K"))
	//	if err != nil {
	//		log.Error().Msgf("Failed to convert env variable ALPHA from string to int: %s", err)
	//	}
	//	log.Debug().Str("Hash", find.hash.String()).Msg("Did not find key")
	//	closest := node.FindKClosest(find.hash, find.senderAddress.ID, k)
	//	data := contact.SerializeContacts(closest)
	//	node.Network.SendFindDataRespMessage(node.ID, find.senderAddress.Address, find.rpcId, &data)
	//}
}

func (find FindValue) ParseOptions(options *[]string) error {
	if (len(*options)) == 0 {
		return errors.New("No hash")
	}

	return nil
}
