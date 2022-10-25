package network

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/network/sender"
	"kademlia/internal/rpc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetwork(t *testing.T) {
	sender, _ := sender.New()
	net := Network{sender}

	senderID := kademliaid.NewRandomKademliaID()
	goodTarget := address.New("127.0.0.1:8888")
	badTarget := address.New("not a real address")

	// Test ping
	err := net.SendPingMessage(senderID, goodTarget)
	assert.NoError(t, err)
	err = net.SendPingMessage(senderID, badTarget)
	assert.Error(t, err)

	// Test pong
	err = net.SendPongMessage(senderID, goodTarget, kademliaid.NewRandomKademliaID())
	assert.NoError(t, err)
	err = net.SendPongMessage(senderID, badTarget, kademliaid.NewRandomKademliaID())
	assert.Error(t, err)

	// Test find_node
	err = net.SendFindContactMessage(&rpc.RPC{Target: goodTarget, RPCId: kademliaid.NewRandomKademliaID()})
	assert.NoError(t, err)
	err = net.SendFindContactMessage(&rpc.RPC{Target: badTarget, RPCId: kademliaid.NewRandomKademliaID()})
	assert.Error(t, err)

	// Test find_node_resp
	content := "someNode"
	err = net.SendFindContactRespMessage(senderID, goodTarget, kademliaid.NewRandomKademliaID(), &content)
	assert.NoError(t, err)
	err = net.SendFindContactRespMessage(senderID, badTarget, kademliaid.NewRandomKademliaID(), &content)
	assert.Error(t, err)

	// Test find_value
	err = net.SendFindDataMessage(&rpc.RPC{Target: goodTarget, RPCId: kademliaid.NewRandomKademliaID()})
	assert.NoError(t, err)
	err = net.SendFindDataMessage(&rpc.RPC{Target: badTarget, RPCId: kademliaid.NewRandomKademliaID()})
	assert.Error(t, err)

	// Test find_value_resp
	content = "someHash"
	err = net.SendFindDataRespMessage(senderID, goodTarget, kademliaid.NewRandomKademliaID(), &content)
	assert.NoError(t, err)
	err = net.SendFindDataRespMessage(senderID, badTarget, kademliaid.NewRandomKademliaID(), &content)
	assert.Error(t, err)

	// Test store
	data := []byte("data")
	err = net.SendStoreMessage(senderID, goodTarget, data)
	assert.NoError(t, err)
	err = net.SendStoreMessage(senderID, badTarget, data)
	assert.Error(t, err)

	// Test refresh
	content = "someHash"
	err = net.SendRefreshMessage(&rpc.RPC{Target: goodTarget, RPCId: kademliaid.NewRandomKademliaID()})
	assert.NoError(t, err)
	err = net.SendRefreshMessage(&rpc.RPC{Target: badTarget, RPCId: kademliaid.NewRandomKademliaID()})
	assert.Error(t, err)
}
