package node_test

import (
	"kademlia/internal/address"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKClosest(t *testing.T) {
	n := node.Node{}
	addr := address.New("127.0.1.1")
	n.Init(addr)
	key := kademliaid.FromString("ffffffffffffffffffffffffffffffffffffffff")
	id1 := kademliaid.FromString("fffffffffffffffffffffffffffffffffffffff0")
	id2 := kademliaid.FromString("ffffffffffffffffffffffffffffffffffffff00")
	id3 := kademliaid.FromString("fffffffffffffffffffffffffffffffffffff000")
	c1 := contact.NewContact(id1, addr)
	c2 := contact.NewContact(id2, addr)
	c3 := contact.NewContact(id3, addr)
	n.RoutingTable.AddContact(c1)
	n.RoutingTable.AddContact(c2)
	n.RoutingTable.AddContact(c3)

	kClosest := n.FindKClosest(key, id1, 3)
	assert.Equal(t, 2, len(kClosest))
}

func TestInitBOOTAndStore(t *testing.T) {
	n := node.Node{}
	addr := address.New("127.0.1.1")
	n.InitBOOT(addr)

	assert.NotNil(t, n)

	n2 := node.Node{}
	addr2 := address.New("127.0.1.1")
	n2.Init(addr2)
	skitt := "0000000000000000000000000000000000000000"
	id := kademliaid.NewKademliaID(&skitt)
	var cc contact.ContactCandidates
	bucketindex := n.RoutingTable.GetBucketIndex(&id)
	bucket := n.RoutingTable.Buckets[bucketindex]
	cc.Append(bucket.GetContactAndCalcDistance(&id))
	n2.Store(&skitt, &cc.Contacts)

	gotten := n2.DataStore.Get(id)

	//cont := n2.RoutingTable.FindClosestContacts(&id, n.ID, 1)

	assert.NotNil(t, gotten)
}

func TestNewRPC(t *testing.T) {
	n := node.Node{}
	addr := address.New("127.0.1.1")
	n.Init(addr)

	rpc := n.NewRPC("hello", addr)

	assert.NotNil(t, rpc)
}

func TestFind(t *testing.T) {
	n := node.Node{}
	addr := address.New("127.0.1.1")
	n.Init(addr)

	skitt := "0000000000000000000000000000000000000000"
	id := kademliaid.NewKademliaID(&skitt)
	c := contact.NewContact(&id, addr)
	n.RoutingTable.AddContact(c)

	skitt2 := "0000000000000000000000000000000000000001"
	id2 := kademliaid.NewKademliaID(&skitt2)
	c2 := contact.NewContact(&id2, addr)
	n.RoutingTable.AddContact(c2)

	skitt3 := "0000000000000000000000000000000000000002"
	id3 := kademliaid.NewKademliaID(&skitt3)
	c3 := contact.NewContact(&id3, addr)
	n.RoutingTable.AddContact(c3)

	skitt4 := "0000000000000000000000000000000000000002"
	id4 := kademliaid.NewKademliaID(&skitt4)
	c4 := contact.NewContact(&id4, addr)
	n.RoutingTable.AddContact(c4)

	res := n.FIND_NODE(&id)

	assert.NotNil(t, res)

	n.Store(&skitt, nil)

	n.FIND_DATA(&id)

	assert.NotNil(t, n.Shortlist.GetData())
}
