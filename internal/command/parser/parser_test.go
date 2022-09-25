package cmdparser_test

import (
	cmdparser "kademlia/internal/command/parser"
	"kademlia/internal/node"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Command interface {
	Execute(node *node.Node) (string, error)

	// Parse the options (i.e. words after command) and set related fields in
	// the struct
	ParseOptions(options []string) error

	PrintUsage() string
}

func TestParseCmd(t *testing.T) {
	var cmd Command

	//should be able to parse a ping command
	// TODO: Should also test that target is set
	cmd = cmdparser.ParseCmd("ping localhost")
	assert.NotNil(t, cmd)

	//should return nil if invalid options were passed
	cmd = cmdparser.ParseCmd("ping") //ping needs a target option
	assert.Nil(t, cmd)

	// should be able to parse a put command
	cmd = cmdparser.ParseCmd("put")
	assert.Nil(t, cmd)
}
