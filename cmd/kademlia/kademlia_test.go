package main_test

import (
	main "kademlia/cmd/kademlia"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	assert.NoError(t, main.InitLogger("1"))
}

func TestGetHostIP(t *testing.T) {
	ip := main.GetHostIP().String()
	assert.NotEqual(t, "", ip)
}
