package forget

import (
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

type Forget struct {
	fileContent string
}

func (forg *Forget) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing forget command")
	key := kademliaid.NewKademliaID(&forg.fileContent)
	node.DataStore.ForgetEntry(key)

	return key.String(), nil
}
