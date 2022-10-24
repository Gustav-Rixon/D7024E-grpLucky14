package refresh_test

import (
	"kademlia/internal/address"
	"kademlia/internal/node"
	"kademlia/internal/rpc/commands/refresh"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptions(t *testing.T) {
	var options []string
	var err error

	adr := address.New("127.0.0.1:1776")
	n := node.Node{}
	n.Init(adr)
	r := refresh.New()
	options = []string{"this is some file content"}
	//fileContent := "this is some file content"
	err = r.ParseOptions(&options)
	assert.NoError(t, err)
	r.Execute(&n)
	//Should never return an error
	assert.NotNil(t, r)
}
