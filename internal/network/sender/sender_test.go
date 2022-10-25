package sender

import (
	"kademlia/internal/address"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSender(t *testing.T) {
	sender, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, sender)

	// Test send to valid address
	err = sender.Send("data", address.New("127.0.0.1:8888"))
	assert.NoError(t, err)

	// Test send to invalid address
	err = sender.Send("data", address.New("not a real address"))
	assert.Error(t, err)
}
