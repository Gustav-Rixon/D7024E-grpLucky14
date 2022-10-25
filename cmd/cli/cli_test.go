package main_test

import (
	main "kademlia/cmd/cli"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrArrayToByteArray(t *testing.T) {
	streng := []string{"1"}
	main.StrArrayToByteArray(streng)
	assert.NotNil(t, main.StrArrayToByteArray(streng))
}
