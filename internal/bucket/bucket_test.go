package bucket_test

import (
	"container/list"
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
	var b *bucket.Bucket
	var bList list.List
	adr := address.New("127.0.0.1")

	// Adding a new contact to a non-full bucket
	// should insert the new contact
	id := kademliaid.NewRandomKademliaID()
	c := contact.NewContact(id, adr)
	b = bucket.NewBucket()
	b.AddContact(c)
	bList = b.GetBucketList()
	assert.Equal(t, bList.Front().Value.(contact.Contact).ID, c.ID)
	assert.Equal(t, bList.Front().Value.(contact.Contact).Address, c.Address)

	// Adding a new contact to a full bucket
	b = bucket.NewBucket()
	for i := 0; i < 20; i++ {
		c = contact.NewContact(kademliaid.NewRandomKademliaID(), adr)
		b.AddContact(c)
	}
	fullBucket := b.GetBucketList()
	assert.Equal(t, fullBucket.Len(), 20)

	// should not add the contact
	newContact := contact.NewContact(kademliaid.NewRandomKademliaID(), adr)
	b.AddContact(newContact)
	newFullBucket := b.GetBucketList()
	assert.Equal(t, newFullBucket.Len(), 20)
	for i, j := newFullBucket.Front(), fullBucket.Front(); i != nil || j != nil; i, j = i.Next(), j.Next() {
		assert.True(t, i.Value.(contact.Contact) == j.Value.(contact.Contact))
	}

	// Adding an already existing contact to the bucket
	// should push the contact to the front of the bucket
	b = bucket.NewBucket()
	var testContact contact.Contact
	for i := 0; i < 20; i++ {
		c = contact.NewContact(kademliaid.NewRandomKademliaID(), adr)
		b.AddContact(c)

		if i == 0 {
			testContact = c
		}
	}
	bList = b.GetBucketList()
	assert.Equal(t, bList.Len(), 20)

	assert.False(t, bList.Front().Value.(contact.Contact) == testContact)
	b.AddContact(testContact)
	bList = b.GetBucketList()
	assert.True(t, bList.Front().Value.(contact.Contact) == testContact)
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
	key := kademliaid.FromString("0000000000000000000000000000000000000001") // random id's cant be trusted
	key3 := kademliaid.FromString("0000000000000000000000000000000000000003")
	inAddr := "127.0.0.1:8888"
	adr := address.New(inAddr)
	b.AddContact(contact.NewContact(key, adr))
	key2 := kademliaid.FromString("0000000000000000000000000000000000000002")
	c := b.GetContactAndCalcDistanceNoRequestor(key2, key3)
	assert.NotNil(t, c)
}
