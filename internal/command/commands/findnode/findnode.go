package findnode

import (
	"errors"
	"fmt"
	"kademlia/internal/constants"
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

	// Get Round 1 contacts
	candidats := node.FindKClosest(tt, node.ID, constants.K)

	fmt.Println("!")
	fmt.Println(candidats)
	fmt.Println("!")
	fmt.Println("!!")
	fmt.Println(candidats[0])
	fmt.Println("!!")
	fmt.Println("!!!")
	fmt.Println(candidats[1])
	fmt.Println("!!!")
	fmt.Println("!!!!")
	fmt.Println(candidats[2])
	fmt.Println("!!!!")

	// Create FIND_NODE RPC's to Round 1 contacts
	rpc := node.NewRPC("FIND_NODE "+id.targetHash, candidats[0].Address)
	rpc1 := node.NewRPC("FIND_NODE "+id.targetHash, candidats[1].Address)
	rpc2 := node.NewRPC("FIND_NODE "+id.targetHash, candidats[2].Address)

	// SEND RPC'S to Round 1 contacts
	node.Network.SendFindContactMessage(&rpc)
	node.Network.SendFindContactMessage(&rpc1)
	node.Network.SendFindContactMessage(&rpc2)

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
