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
	//Should never return an error
	assert.NoError(t, p.ParseOptions(&options))
	options = []string{}
	assert.Error(t, p.ParseOptions(&options))
}
