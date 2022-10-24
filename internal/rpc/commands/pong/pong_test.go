package pong_test

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"kademlia/internal/rpc/commands/pong"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptions(t *testing.T) {
	adr := address.New("127.0.0.1:1776")
	p := pong.New(kademliaid.NewRandomKademliaID(), adr, kademliaid.NewRandomKademliaID())
	options := []string{"hello", "abc"}
	n := node.Node{}
	addr := address.New("127.0.1.1")
	n.Init(addr)
	p.Execute(&n)
	//Should never return an error
	assert.NoError(t, p.ParseOptions(&options))
}
