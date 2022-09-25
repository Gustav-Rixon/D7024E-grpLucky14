package datastore

import (
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/rpc"
)

type DataMap = map[kademliaid.KademliaID]*Data

type DataStore struct {
	store DataMap
}

type Data struct {
	value    string
	Contacts *[]contact.Contact
}

// Create the hash map
func New() DataStore {
	return DataStore{make(DataMap)}
}

// Insert a data into the store.
func (datastorage *DataStore) Insert(value string, contacts *[]contact.Contact, sender rpc.Sender) {
	id := kademliaid.NewKademliaID(&value)
	data := Data{}
	data.value = value
	data.Contacts = contacts
	datastorage.store[id] = &data
}

// Gets the value from the store associated with the key.
// Returns an empty string if the key is not found
func (datastorage *DataStore) GetValue(key kademliaid.KademliaID) string {
	data := datastorage.store[key]
	if data != nil {
		return data.value

	}
	return ""
}

// Gets the value from the store associated with the key.
// Returns an empty string if the key is not found
func (d *DataStore) Get(key kademliaid.KademliaID) string {
	data := d.store[key]
	if data != nil {
		return data.value

	}
	return ""

}
