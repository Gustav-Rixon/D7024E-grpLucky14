package add_test

import (
	"kademlia/internal/address"
	"kademlia/internal/command/commands/add"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptions(t *testing.T) {
	var addCmd *add.Add
	var err error

	addr := address.New("127.0.0.1:1234")
	n := node.Node{}

	n.Init(addr)

	skitt := "0000000000000000000000000000000000000000"
	id := kademliaid.NewKademliaID(&skitt)
	c1 := contact.NewContact(&id, addr)
	n.RoutingTable.AddContact(c1)

	skitt2 := "0000000000000000000000000000000000000001"
	id2 := kademliaid.NewKademliaID(&skitt2)
	c2 := contact.NewContact(&id2, addr)
	n.RoutingTable.AddContact(c2)

	skitt3 := "0000000000000000000000000000000000000002"
	id3 := kademliaid.NewKademliaID(&skitt3)
	c3 := contact.NewContact(&id3, addr)
	n.RoutingTable.AddContact(c3)

	skitt4 := "0000000000000000000000000000000000000002"
	id4 := kademliaid.NewKademliaID(&skitt4)
	c4 := contact.NewContact(&id4, addr)
	n.RoutingTable.AddContact(c4)

	n.FIND_NODE(&id4)

	// should not return an error if content is specified
	addCmd = new(add.Add)
	err = addCmd.ParseOptions([]string{"127.0.0.1:1234", "content"})
	assert.Nil(t, err)

	addCmd.Execute(&n)
}

func TestPrintUsage(t *testing.T) {
	// should be equal
	var addCmd *add.Add
	assert.Equal(t, addCmd.PrintUsage(), "Usage: addcontact {nodeID} {address}")

}
