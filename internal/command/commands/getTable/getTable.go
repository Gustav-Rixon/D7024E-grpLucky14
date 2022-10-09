package getTable

import (
	"fmt"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

type GetTable struct {
	Target string
}

func (p GetTable) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing getcontacts command")
	fmt.Println(node.RoutingTable.GetContacts()) //For testing
	return node.RoutingTable.GetContacts(), nil
}

func (g *GetTable) ParseOptions(options []string) error {
	return nil
}

func (g *GetTable) PrintUsage() string {
	return "Usage: getcontacts"
}
