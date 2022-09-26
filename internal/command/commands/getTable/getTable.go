package getTable

import (
	"errors"
	"fmt"
	"kademlia/internal/node"
)

type GetTable struct {
	Target string
}

func (p GetTable) Execute(node *node.Node) (string, error) {
	fmt.Println(node.RoutingTable.GetContacts())
	return "Table its broken!", nil
}

func (p *GetTable) ParseOptions(options []string) error {
	if len(options) < 1 {
		return errors.New("Missing Target")
	}
	p.Target = options[0]
	return nil
}

func (p *GetTable) PrintUsage() string {
	return "Usage: ping {target address}"
}
