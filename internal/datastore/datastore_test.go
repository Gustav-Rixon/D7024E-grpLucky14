package datastore_test

import (
	"testing"

	"kademlia/internal/contact"
	"kademlia/internal/datastore"
	"kademlia/internal/kademliaid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SenderMock struct {
	mock.Mock
}

func TestGetValue(t *testing.T) {
	var d datastore.DataStore

	// Should be able to  get
	d = datastore.New()
	value := "hello"
	contacts := &[]contact.Contact{}
	d.Insert(value, contacts, nil, true)
	assert.Equal(t, d.GetValue(kademliaid.NewKademliaID(&value)), "hello")
	assert.Equal(t, d.Get(kademliaid.NewKademliaID(&value)), "hello")

	// Should not be able to get non-existent key
	d = datastore.New()
	value = "hello"
	assert.Equal(t, d.GetValue(kademliaid.NewKademliaID(&value)), "")
	assert.Equal(t, d.Get(kademliaid.NewKademliaID(&value)), "")
}

func TestInsert(t *testing.T) {
	var d datastore.DataStore
	var contacts *[]contact.Contact
	value := "hello"

	//should be able to insert
	d = datastore.New()
	contacts = &[]contact.Contact{}
	d.Insert(value, contacts, nil, true)
	assert.Equal(t, d.GetValue(kademliaid.NewKademliaID(&value)), "hello")

}

func TestRefresh(t *testing.T) {
	var d datastore.DataStore

	// Should be able to  get
	d = datastore.New()
	value := "hello"
	contacts := &[]contact.Contact{}
	d.Insert(value, contacts, nil, true)
	assert.Equal(t, d.Refresh(kademliaid.NewKademliaID(&value)), "cool")

	// Should not be able to get non-existent key
	d = datastore.New()
	value = "hello"
	assert.Equal(t, d.Refresh(kademliaid.NewKademliaID(&value)), "")
}

func TestGetForget(t *testing.T) {
	var d datastore.DataStore

	// Should be able to  get
	d = datastore.New()
	value := "hello"
	contacts := &[]contact.Contact{}
	d.Insert(value, contacts, nil, false)
	assert.Equal(t, d.GetForget(kademliaid.NewKademliaID(&value)), false)

	// Should not be able to get non-existent key
	d = datastore.New()
	value = "hello"
	assert.Equal(t, d.GetForget(kademliaid.NewKademliaID(&value)), true)
}

func TestSetForget(t *testing.T) {
	var d datastore.DataStore

	// Should be able to  get
	d = datastore.New()
	value := "hello"
	contacts := &[]contact.Contact{}
	d.Insert(value, contacts, nil, false)
	assert.Equal(t, d.SetForget(kademliaid.NewKademliaID(&value), false), "cool")

	d = datastore.New()

	assert.Equal(t, d.SetForget(kademliaid.NewKademliaID(&value), true), "")

}
