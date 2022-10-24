package findnodeRPC_test

import (
	"kademlia/internal/address"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"kademlia/internal/rpc/commands/findnodeRPC"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptions(t *testing.T) {
	adr := address.New("127.0.0.1:1776")
	c := contact.NewContact(kademliaid.NewRandomKademliaID(), adr)
	p := findnodeRPC.New(&c, kademliaid.NewRandomKademliaID())
	options := []string{"hello", "abc"}
	n := node.Node{}
	addr := address.New("127.0.1.1")
	n.Init(addr)
	//Should never return an error
	assert.NoError(t, p.ParseOptions(&options))
	options = []string{}
	assert.Error(t, p.ParseOptions(&options))
}
