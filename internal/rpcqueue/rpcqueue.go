//https://www.rabbitmq.com/tutorials/tutorial-six-go.html

package rpcqueue

import (
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"sync"
)

type Entry struct {
	Channel chan string
	RPCID   *kademliaid.KademliaID
	Contact *contact.Contact
}

type RPCQueue struct {
	lock    sync.Mutex //The queue should be protected by a mutex lock
	content map[kademliaid.KademliaID]*Entry
}

// Create a new RPCQueue
func New() *RPCQueue {
	return &RPCQueue{
		content: make(map[kademliaid.KademliaID]*Entry),
	}
}

// Add RPC to queue
func (queue *RPCQueue) Add(RPCID *kademliaid.KademliaID, contact *contact.Contact) {
	queue.content[*RPCID] = &Entry{RPCID: RPCID, Contact: contact, Channel: make(chan string)}
}

// Delete RPC from queue
func (queue *RPCQueue) Delete(RPCID *kademliaid.KademliaID) {
	delete(queue.content, *RPCID)
}

func (pool *RPCQueue) WithLock(f func()) {
	pool.lock.Lock()
	f()
	pool.lock.Unlock()
}

func (pool *RPCQueue) GetEntry(rpcId *kademliaid.KademliaID) *Entry {
	return pool.content[*rpcId]
}
