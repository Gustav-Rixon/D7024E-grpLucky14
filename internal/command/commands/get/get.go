package get

//TODO GET THE BREaD
import (
	"errors"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

type Get struct {
	hash kademliaid.KademliaID
}

func (get *Get) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing get command")

	// Check local storage
	value := node.DataStore.Get(get.hash)
	if value != "" {
		value += ", from local node"
	}
	//TODO FIND VAULE ON A DIFFRENT NODE
	if value == "" {
		return "", errors.New("Key not found")
	}
	return value, nil
}

func (get *Get) ParseOptions(options []string) error {
	if len(options) < 1 {
		return errors.New("Missing hash")
	}
	get.hash = *kademliaid.FromString(options[0])
	return nil
}

func (get *Get) PrintUsage() string {
	return "USAGE: get <hash>"
}
