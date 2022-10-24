package findnoderesponse_test

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"kademlia/internal/rpc/commands/findnoderesponse"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptions(t *testing.T) {
	addr := address.New("127.0.0.1:1776")
	//c := contact.NewContact(kademliaid.NewRandomKademliaID(), adr)
	p := findnoderesponse.New(kademliaid.NewRandomKademliaID())
	options := []string{"hello", "abc"}
	n := node.Node{}

	n.Init(addr)

	assert.NotNil(t, findnoderesponse.DeserializeContacts("0 1 2 3 4 5 6", kademliaid.NewRandomKademliaID()))
	//Should never return an error
	assert.NoError(t, p.ParseOptions(&options))
	options = []string{}
	assert.Error(t, p.ParseOptions(&options))
}
