package nodeLookup

import (
	"errors"
	"fmt"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

type NodeLookup struct {
}

func (look *NodeLookup) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing TEST")

	candis := node.NodeLookup(node.ID)
	fmt.Println(candis)

	return "candis", nil
}

func (look *NodeLookup) ParseOptions(options []string) error {
	if len(options) < 1 {
		return errors.New("Missing Life")
	}
	return nil
}

func (look *NodeLookup) PrintUsage() string {
	return "USAGE: get <hash>"
}
