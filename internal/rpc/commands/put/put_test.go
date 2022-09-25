package put_test

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"kademlia/internal/rpc/commands/put"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	var s put.Put
	var options []string
	var err error

	// Should set file content if passed
	options = []string{"this is some file content"}
	fileContent := "this is some file content"
	err = s.ParseOptions(&options)
	assert.NoError(t, err)
	node := node.Node{}
	node.Init(address.New(""))
	s.Execute(&node)
	assert.Equal(t, fileContent, node.DataStore.GetValue(kademliaid.NewKademliaID(&fileContent)))
}
