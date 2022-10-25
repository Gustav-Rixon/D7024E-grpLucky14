package main_test

import (
	main "kademlia/cmd/kademlia"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	assert.NoError(t, main.InitLogger("1"))
}

func TestBootstrap(t *testing.T) {
	//no this is cursed
}
