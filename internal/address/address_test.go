package address_test

import (
	"kademlia/internal/address"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	// should return the address as a string
	inAddr := "127.0.0.1:8888"
	adr := address.New(inAddr)
	assert.Equal(t, adr.String(), inAddr)
}
