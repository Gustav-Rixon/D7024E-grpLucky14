package findvalueRPC_test

import (
	"kademlia/internal/address"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"kademlia/internal/rpc/commands/findvalueRPC"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptions(t *testing.T) {
	addr := address.New("127.0.0.1:1776")
	c := contact.NewContact(kademliaid.NewRandomKademliaID(), addr)
	p := findvalueRPC.New(&c, kademliaid.NewRandomKademliaID())
	options := []string{"hello", "abc"}
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

	//Should never return an error
	assert.NoError(t, p.ParseOptions(&options))

	p.Execute(&n)

	options = []string{}
	assert.Error(t, p.ParseOptions(&options))
}
