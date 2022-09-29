package findnoderesponse

import (
	"errors"
	"fmt"
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

func (Resp *FindNodeResponse) Execute(node *node.Node) {
	log.Trace().Msg("Executing FIND_NODE_RESP RPC")

	fmt.Println("!!!!!!!")
	fmt.Println(Resp)
	fmt.Println(Resp.data)
	fmt.Println(*Resp.data)
	fmt.Println(&Resp.data)
	fmt.Println("!!!!!!!")

}

func (Resp *FindNodeResponse) ParseOptions(options *[]string) error {
	if len(*options) == 0 {
		return errors.New("Empty RESP")
	}
	data := strings.Join(*options, "")
	Resp.data = &data
	return nil
}
