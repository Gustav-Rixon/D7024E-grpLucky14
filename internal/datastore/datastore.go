package datastore

import (
	"fmt"
	"kademlia/internal/constants"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/rpc"
	"sync"
	"time"
)

type DataMap = map[kademliaid.KademliaID]*Data

type DataStore struct {
	store   DataMap
	timeTTL time.Duration
	lock    *sync.RWMutex
}

type Data struct {
	expiry   *time.Time
	value    string
	Contacts *[]contact.Contact
	forget   bool
}

// TTL functionality inspired by: https://github.com/Konstantin8105/SimpleTTL/blob/master/simplettl.go
const TTL time.Duration = constants.DataTTL * time.Second

// Create the hash map
func New() DataStore {

	datastore := &DataStore{
		store:   make(DataMap),
		timeTTL: TTL,
		lock:    &sync.RWMutex{},
	}

	go func() {
		ticker := time.NewTicker(datastore.timeTTL)
		for {
			now := <-ticker.C

			datastore.lock.Lock()

			for id, entry := range datastore.store {
				if entry.expiry != nil && entry.expiry.Before(now) {
					fmt.Println(datastore.store[id].value)
					delete(datastore.store, id)
					fmt.Println("Data object expired")
				}
			}
			datastore.lock.Unlock()
		}
	}()
	return *datastore
}

// Insert a data into the store.
func (datastorage *DataStore) Insert(value string, contacts *[]contact.Contact, sender rpc.Sender, forg bool) {
	datastorage.lock.Lock()
	defer datastorage.lock.Unlock()
	id := kademliaid.NewKademliaID(&value)
	data := Data{}
	data.value = value
	data.Contacts = contacts
	expiry := time.Now().Add(TTL)
	data.expiry = &expiry
	data.forget = forg
	datastorage.store[id] = &data
}

// Gets the value from the store associated with the key.
// Returns an empty string if the key is not found
func (datastorage *DataStore) GetValue(key kademliaid.KademliaID) string {
	datastorage.lock.Lock()
	defer datastorage.lock.Unlock()
	data := datastorage.store[key]

	if data != nil {
		expiry := time.Now().Add(TTL)
		data.expiry = &expiry
		return data.value

	}
	return ""
}

// Gets the value from the store associated with the key.
// Returns an empty string if the key is not found
func (datastorage *DataStore) Get(key kademliaid.KademliaID) string {
	datastorage.lock.Lock()
	defer datastorage.lock.Unlock()
	data := datastorage.store[key]

	if data != nil {
		expiry := time.Now().Add(TTL)
		data.expiry = &expiry
		return data.value
	}
	return ""

}

func (datastorage *DataStore) Refresh(key kademliaid.KademliaID) string {
	datastorage.lock.Lock()
	defer datastorage.lock.Unlock()

	data := datastorage.store[key]

	if data != nil {
		expiry := time.Now().Add(TTL)
		data.expiry = &expiry
		return "cool"
	}
	return ""
}

func (datastorage *DataStore) GetForget(key kademliaid.KademliaID) bool {
	datastorage.lock.Lock()
	defer datastorage.lock.Unlock()

	data := datastorage.store[key]
	if data != nil {
		return data.forget
	}
	return true
}

func (datastorage *DataStore) SetForget(key kademliaid.KademliaID, forg bool) string {
	datastorage.lock.Lock()
	defer datastorage.lock.Unlock()

	data := datastorage.store[key]
	if data != nil {
		data.forget = forg
		return "cool"
	}
	return ""
}
