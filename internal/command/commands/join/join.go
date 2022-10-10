package join

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

type Join struct {
	Target string
}

func (p *Join) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing join command")
	adr := address.New(p.Target)
	node.Network.SendPingMessage(node.ID, adr)
	//node.FIND_NODE(adr) //always ask route node when joining
	node.FIND_NODE(kademliaid.FromString(p.Target))
	return "Joined network on known node", nil
}

func (p *Join) ParseOptions(options []string) error {
	p.Target = options[0]
	return nil
}

func (j *Join) PrintUsage() string {
	return "Usage: join"
}
