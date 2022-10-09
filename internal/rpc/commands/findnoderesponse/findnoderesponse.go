package findnoderesponse

import (
	"errors"
	"fmt"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"strings"

	"github.com/rs/zerolog/log"
)

type FindNodeResponse struct {
	rpcId *kademliaid.KademliaID
	data  *string
}

func New(rpcId *kademliaid.KademliaID) *FindNodeResponse {
	return &FindNodeResponse{rpcId: rpcId}
}

// Will insert values to routing table
func (Resp *FindNodeResponse) Execute(node *node.Node) {
	log.Trace().Msg("Executing FIND_NODE_RESP RPC")

	test := *Resp.data

	test2, err := contact.Deserialize(&test)
	node.RoutingTable.AddContact()
	fmt.Println(err)
	fmt.Println("!!!!")
	fmt.Println(test)
	fmt.Println("!!!!")
	fmt.Println(test2)

}

func (Resp *FindNodeResponse) ParseOptions(options *[]string) error {
	if len(*options) == 0 {
		return errors.New("Empty RESP")
	}
	data := strings.Join(*options, "")
	Resp.data = &data
	return nil
}
