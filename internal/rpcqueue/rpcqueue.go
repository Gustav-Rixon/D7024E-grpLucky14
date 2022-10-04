//https://www.rabbitmq.com/tutorials/tutorial-six-go.html

package rpcqueue

import (
	"kademlia/internal/rpc"
	"sync"
)

var rpcQueue []rpc.RPC     // The queue for holding all RPC's
var queueMutex sync.Mutex  // Mutex lock for modifying the queue
var execSem sync.WaitGroup // Semaphore for counting all running execution threads

// Adds an RPC request to the execution queue, begins running through the queue if semaphore count = 0
func addToQueue() {

}

func handleLoop() {
	queueMutex.Lock()

}
