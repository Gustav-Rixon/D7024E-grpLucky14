package findnode

import (
	"errors"
	"fmt"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"strings"

	"github.com/rs/zerolog/log"
)

type FindNode struct {
	targetHash string
}

func (id *FindNode) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing find node")

	tt := kademliaid.FromString(id.targetHash)

	candidats := node.FindKClosest(tt, node.ID, 3)

	fmt.Println("!!!!!!!!!!!!!!")
	fmt.Println(id.targetHash)
	fmt.Println("!!!!!!!!!!!!!!")

	rpc := node.NewRPC("FIND_NODE "+id.targetHash, candidats[0].Address)

	// SEND RPC'S?
	node.Network.SendFindContactMessage(&rpc)

	return "contact", nil
}

func (id *FindNode) ParseOptions(options []string) error {
	if len(options) < 1 {
		return errors.New("Missing hash")
	}
	id.targetHash = strings.Join(options[0:], " ")
	return nil
}

func (look *FindNode) PrintUsage() string {
	return "USAGE: get <hash>"
}
