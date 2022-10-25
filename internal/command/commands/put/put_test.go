package put_test

import (
	"kademlia/internal/address"
	"kademlia/internal/command/commands/put"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptions(t *testing.T) {
	var putCmd *put.Put
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
	putCmd = new(put.Put)
	err = putCmd.ParseOptions([]string{"address", "content"})
	assert.Nil(t, err)

	putCmd.Execute(&n)

	// should return an error
	putCmd = new(put.Put)
	err = putCmd.ParseOptions([]string{})
	assert.NotNil(t, err)

}

func TestPrintUsage(t *testing.T) {
	// should be equal
	var putCmd *put.Put
	assert.Equal(t, putCmd.PrintUsage(), "put <file content>")

}
