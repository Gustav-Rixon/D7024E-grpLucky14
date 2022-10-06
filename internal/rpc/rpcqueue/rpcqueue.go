//https://www.rabbitmq.com/tutorials/tutorial-six-go.html

package rpcqueue

import (
	"kademlia/internal/node"
	"kademlia/internal/rpc/rpcparser"
	"sync"
)

type RPCQueue struct {
	queue      []rpcparser.RPCCommand // The node's RPC command queue
	queueMutex *sync.Mutex            // Mutex lock for modifying the queue
	loopMutex  *sync.Mutex            // Mutex lock for handle loop
	loopActive bool
}

func New() RPCQueue {
	q := RPCQueue{
		loopActive: false,
		loopMutex:  &sync.Mutex{},
		queueMutex: &sync.Mutex{},
	}

	return q
}

// Adds an RPC request to the RPC queue and signals the handle loop to start
func (q *RPCQueue) AddToQueue(thisNode *node.Node, rpc rpcparser.RPCCommand) {
	q.queueMutex.Lock()
	q.queue = append(q.queue, rpc)
	q.queueMutex.Unlock()

	go q.initHandleLoop(thisNode)
}

// Idempotent call to start a handle loop if there is no currently running loop
func (q *RPCQueue) initHandleLoop(thisNode *node.Node) {
	q.loopMutex.Lock()
	if q.loopActive {
		q.loopMutex.Unlock()
		return
	} else {
		q.loopActive = true
		q.loopMutex.Unlock()
		q.handleLoop(thisNode)
	}
}

// Loop that executes all RPC's in the queue in the order they have arrived, until there are no more RPC's to be handled
func (q *RPCQueue) handleLoop(thisNode *node.Node) {
	for {
		q.loopMutex.Lock()
		q.queueMutex.Lock()

		// Check if RPC queue contains any RPC's
		if len(q.queue) > 0 {
			// Pop head RPC
			firstRPC := q.queue[0]
			q.queue = q.queue[1:]

			q.loopMutex.Unlock()
			q.queueMutex.Unlock()

			// Execute head RPC
			firstRPC.Execute(thisNode)

		} else { // If RPCQueue is empty: set loop boolean to false and exit loop
			q.loopActive = false

			q.loopMutex.Unlock()
			q.queueMutex.Unlock()
			break
		}
	}
}
