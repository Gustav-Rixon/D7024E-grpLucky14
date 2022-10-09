package store

import (
	"errors"
	"kademlia/internal/node"

	"strings"

	"github.com/rs/zerolog/log"
)

type Store struct {
	Content string //For abstraction sake this will always be a string right?
}

func (store *Store) Execute(node *node.Node) {
	log.Trace().Msg("Executing FIND_NODE RPC")
	node.Store(&store.Content, nil)
}

func (store *Store) ParseOptions(options *[]string) error {
	if len(*options) == 0 {
		return errors.New("Received empty STORE RPC")
	}
	store.Content = strings.Join(*options, " ")
	return nil
}
