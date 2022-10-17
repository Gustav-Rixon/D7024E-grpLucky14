package bucket_test

import (
	"kademlia/internal/address"
	"kademlia/internal/bucket"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBucket(t *testing.T) {
	// should create a new and empty bucket
	b := bucket.NewBucket()
	assert.NotNil(t, b)
	assert.Equal(t, b.Len(), 0)
}

func TestAddContact(t *testing.T) {
	b := bucket.NewBucket()
	key := kademliaid.NewRandomKademliaID()
	inAddr := "127.0.0.1:8888"
	adr := address.New(inAddr)
	b.AddContact(contact.NewContact(key, adr))
	assert.NotNil(t, b)
	assert.Equal(t, b.Len(), 1)
}

func TestGetContactAndCalcDistance(t *testing.T) {
	b := bucket.NewBucket()
	key := kademliaid.NewRandomKademliaID()
	inAddr := "127.0.0.1:8888"
	adr := address.New(inAddr)
	b.AddContact(contact.NewContact(key, adr))
	key2 := kademliaid.NewRandomKademliaID()
	c := b.GetContactAndCalcDistance(key2)
	assert.NotNil(t, c)
}

func TestGetContactAndCalcDistanceNoRequestor(t *testing.T) {
	b := bucket.NewBucket()
	key := kademliaid.NewRandomKademliaID()
	inAddr := "127.0.0.1:8888"
	adr := address.New(inAddr)
	b.AddContact(contact.NewContact(key, adr))
	key2 := kademliaid.NewRandomKademliaID()
	c := b.GetContactAndCalcDistanceNoRequestor(key2, key)
	assert.NotNil(t, c)
}
