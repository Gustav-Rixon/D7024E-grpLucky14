package findvalueresp_test

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	findvalueresp "kademlia/internal/rpc/commands/findvalueresponse"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptions(t *testing.T) {
	addr := address.New("127.0.0.1:1776")
	//c := contact.NewContact(kademliaid.NewRandomKademliaID(), addr)
	p := findvalueresp.New(kademliaid.NewRandomKademliaID(), kademliaid.NewRandomKademliaID())
	options := []string{"hello", "abc"}
	n := node.Node{}

	n.Init(addr)
	assert.NotNil(t, findvalueresp.DeserializeContacts("0 1 2 3 4 5 6", kademliaid.NewRandomKademliaID()))
	//Should never return an error
	assert.NoError(t, p.ParseOptions(&options))
	options = []string{}
	assert.Error(t, p.ParseOptions(&options))
}
