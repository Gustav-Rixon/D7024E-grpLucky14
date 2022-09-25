package put_test

import (
	"kademlia/internal/command/commands/put"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptions(t *testing.T) {
	var putCmd *put.Put
	var err error

	// should not return an error if content is specified
	putCmd = new(put.Put)
	err = putCmd.ParseOptions([]string{"address", "content"})
	assert.Nil(t, err)

	// should return an error
	putCmd = new(put.Put)
	err = putCmd.ParseOptions([]string{})
	assert.NotNil(t, err)

}

func TestPrintUsage(t *testing.T) {
	// should be equal
	var putCmd *put.Put
	assert.Equal(t, putCmd.PrintUsage(), "put <file content>")

}
