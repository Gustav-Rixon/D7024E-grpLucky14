package refresh_test

import (
	"kademlia/internal/address"
	"kademlia/internal/node"
	"kademlia/internal/rpc/commands/refresh"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptions(t *testing.T) {
	adr := address.New("127.0.0.1:1776")
	n := node.Node{}
	n.Init(adr)
	r := refresh.New()
	var str []string
	str = append(str, "REFRESH 0")
	//Should never return an error
	assert.NotNil(t, r)
	assert.NoError(t, r.ParseOptions(&str))
}
