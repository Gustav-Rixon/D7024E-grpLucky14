//https://www.rabbitmq.com/tutorials/tutorial-six-go.html

package rpcqueue

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"kademlia/internal/rpc/commands/ping"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	node := node.Node{}
	node.Init(address.New("192.0.0.1:8888"))

	q := New()
	q.AddToQueue(&node, ping.New(kademliaid.NewRandomKademliaID(), address.New("127.0.0.1:8888"), kademliaid.NewRandomKademliaID()))
	q.AddToQueue(&node, ping.New(kademliaid.NewRandomKademliaID(), address.New("127.0.0.1:8888"), kademliaid.NewRandomKademliaID()))
	q.AddToQueue(&node, ping.New(kademliaid.NewRandomKademliaID(), address.New("127.0.0.1:8888"), kademliaid.NewRandomKademliaID()))
	time.Sleep(time.Millisecond * 5)

	// The queue should be handled after 5 milliseconds
	assert.Empty(t, q.queue)
}
