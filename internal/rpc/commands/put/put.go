package put

import (
	"errors"
	"kademlia/internal/node"
	"strings"

	"github.com/rs/zerolog/log"
)

type Put struct {
	Content string //For abstraction sake this will always be a string right?
}

func (put *Put) Execute(node *node.Node) {
	log.Trace().Msg("Executing FIND_NODE RPC")
	node.Store(&put.Content, nil)
}

func (put *Put) ParseOptions(options *[]string) error {
	if len(*options) == 0 {
		return errors.New("Received empty STORE RPC")
	}
	put.Content = strings.Join(*options, " ")
	return nil
}
