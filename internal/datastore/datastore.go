package datastore

import (
	"kademlia/internal/kademliaid"
)

type DataMap = map[kademliaid.KademliaID]*Data

type DataStore struct {
	store DataMap
}

type Data struct {
	value string
}

func New() DataStore {
	return DataStore{make(DataMap)}
}

// Insert a data into the store.
func (datastorage *DataStore) Insert(value string) {
	id := kademliaid.NewKademliaIDTEST(&value)
	data := Data{}
	data.value = value
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
