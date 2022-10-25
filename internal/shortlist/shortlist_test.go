package shortlist

import (
	"kademlia/internal/address"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortlist(t *testing.T) {
	// Init contacts
	contacts := []contact.Contact{contact.NewContact(kademliaid.FromString("0000000000000000000000000000000000000002"), address.New("127.0.0.2:8888")),
		contact.NewContact(kademliaid.FromString("0000000000000000000000000000000000000003"), address.New("127.0.0.3:8888")),
		contact.NewContact(kademliaid.FromString("0000000000000000000000000000000000000004"), address.New("127.0.0.4:8888"))}

	for i := range contacts {
		contacts[i].CalcDistance(kademliaid.FromString("0000000000000000000000000000000000000001"))
	}

	// Init shortlist
	sl := NewShortlist(kademliaid.FromString("0000000000000000000000000000000000000001"), contacts)

	// Make sure less returns true when j = nil and false when i = nil
	assert.True(t, sl.Less(0, 3))
	assert.False(t, sl.Less(3, 0))

	// Attempt adding duplicate
	sl.Add(&contacts[0])

	// Add another contact
	c1 := contact.NewContact(kademliaid.FromString("0000000000000000000000000000000000000005"), address.New("127.0.0.5:8888"))
	c1.CalcDistance(kademliaid.FromString("0000000000000000000000000000000000000001"))
	sl.Add(&c1)

	// Add shortlist's desired target
	c1 = contact.NewContact(kademliaid.FromString("0000000000000000000000000000000000000001"), address.New("127.0.0.5:8888"))
	c1.CalcDistance(kademliaid.FromString("0000000000000000000000000000000000000001"))
	sl.Add(&c1)

	// Assert first in shortlist is target
	assert.Equal(t, c1, sl.GetContacts()[0])

	// Test data
	sl.AddFoundData(kademliaid.NewRandomKademliaID(), "data")
	assert.NotNil(t, sl.GetData())
	assert.NotNil(t, sl.GetDataHost())
}
