package refresh

import (
	"errors"
	"fmt"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

type RefreshRPC struct {
	hash *string
}

func New() *RefreshRPC {
	return &RefreshRPC{}
}

func (refr *RefreshRPC) Execute(node *node.Node) {
	log.Trace().Msg("Executing Refresh RPC")

	key := kademliaid.FromString(*refr.hash)

	if value := node.DataStore.Refresh(*key); value != "" {
		fmt.Println("Object has been found and refreshed")
	} else {
		fmt.Println("ERROR object wasn't found and couldn't be refreshed, something is wrong smh")
	}
}

func (refr *RefreshRPC) ParseOptions(options *[]string) error {
	if (len(*options)) == 0 {
		return errors.New("No hash")
	}

	refr.hash = &(*options)[0]
	return nil
}
