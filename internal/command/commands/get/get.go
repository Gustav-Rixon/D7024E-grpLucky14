package get

//TODO GET THE BREaD
import (
	"errors"
	"fmt"
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
	fmt.Println(value)
	key := kademliaid.NewKademliaID(&value)
	if value != "" {
		value += ", from local node"
	} else {

		// GOGOG RPC FIND_VALUE
		//RPC DID NOT FIND
		//RPC FOUND
		closestNodes := node.FindKClosest(&key, nil, 3)
		targetNode := closestNodes[0]
		value += ", from " + targetNode.String()

	}
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
