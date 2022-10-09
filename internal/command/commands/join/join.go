package join

import (
	"errors"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"strings"

	"github.com/rs/zerolog/log"
)

type Join struct {
	LookingUp string
}

func (j *Join) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing join command")
	tt := kademliaid.FromString(j.LookingUp)
	node.FIND_NODE(tt)
	return "Joined network on known node", nil
}

func (j *Join) ParseOptions(options []string) error {
	if len(options) < 1 {
		return errors.New("Missing target")
	}
	j.LookingUp = strings.Join(options[0:], " ")
	return nil
}

func (j *Join) PrintUsage() string {
	return "Usage: join"
}
