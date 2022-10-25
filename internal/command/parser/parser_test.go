package cmdparser_test

import (
<<<<<<< Updated upstream
=======
	"kademlia/internal/command/commands/add"
	"kademlia/internal/command/commands/forget"
	"kademlia/internal/command/commands/get"
	"kademlia/internal/command/commands/getTable"
	"kademlia/internal/command/commands/getid"
	"kademlia/internal/command/commands/join"
	"kademlia/internal/command/commands/ping"
	"kademlia/internal/command/commands/put"
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
	cmd = cmdparser.ParseCmd("put")
	assert.Nil(t, cmd)
=======
	cmd = cmdparser.ParseCmd("put content")
	assert.Equal(t, reflect.TypeOf(&put.Put{}), reflect.TypeOf(cmd))

	// should be able to parse a get command
	cmd = cmdparser.ParseCmd("get contentHash")
	assert.Equal(t, reflect.TypeOf(&get.Get{}), reflect.TypeOf(cmd))

	// should also be able to parse all other types of commands
	cmd = cmdparser.ParseCmd("add id address")
	assert.Equal(t, reflect.TypeOf(&add.Add{}), reflect.TypeOf(cmd))

	cmd = cmdparser.ParseCmd("getTable")
	assert.Equal(t, reflect.TypeOf(&getTable.GetTable{}), reflect.TypeOf(cmd))

	cmd = cmdparser.ParseCmd("getid")
	assert.Equal(t, reflect.TypeOf(&getid.GetId{}), reflect.TypeOf(cmd))

	cmd = cmdparser.ParseCmd("join target")
	assert.Equal(t, reflect.TypeOf(&join.Join{}), reflect.TypeOf(cmd))

	cmd = cmdparser.ParseCmd("forget content")
	assert.Equal(t, reflect.TypeOf(&forget.Forget{}), reflect.TypeOf(cmd))
>>>>>>> Stashed changes
}
