package add

import (
	"errors"
	"fmt"
	"kademlia/internal/address"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
)

type Add struct {
	Id, Address string
}

func (adr *Add) Execute(node *node.Node) (string, error) {
	addr := address.New(adr.Address)
	node.RoutingTable.AddContact(contact.NewContact(kademliaid.FromString(adr.Id), addr))
	return "Contact added to node's routingTable: " + fmt.Sprint(addr.String()), nil
}

func (adr *Add) ParseOptions(options []string) error {
	if len(options) < 2 {
		return errors.New("Missing contact id or address")
	}

	adr.Id = options[0]
	adr.Address = options[1]
	return nil
}

func (a *Add) PrintUsage() string {
	return "Usage: addcontact {nodeID} {address}"
}
