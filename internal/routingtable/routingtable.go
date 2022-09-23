package routingtable

import (
	"kademlia/internal/bucket"
	"kademlia/internal/kademliaid"
	. "kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"net"
	"sort"
)

const bucketSize = 20

// Routing table used for testing
var rt RoutingTable

// Definitely move this
func GetRT() RoutingTable {
	return rt
}

// RoutingTable definition
// keeps a refrence contact of me and an array of buckets
type RoutingTable struct {
	Me      node.Node
	buckets [kademliaid.IDLength * 8]*bucket.Bucket
}

// NewRoutingTable returns a new instance of a RoutingTable
func NewRoutingTable(me node.Node) {
	//rt := &RoutingTable{}
	for i := 0; i < kademliaid.IDLength*8; i++ {
		rt.buckets[i] = bucket.NewBucket()
	}
	rt.Me = me
}

// AddContact add a new contact to the correct Bucket
// ********************************************************************
// THIS PROBABLY ADDS CONTACTS TO THE WRONG BUCKETS, getBucketIndex() should take the node that we intend to put in, right now its taking itself????
func (routingTable *RoutingTable) AddContact(id [kademliaid.IDLength]byte, ip net.IP) {
	bucketIndex := routingTable.getBucketIndex(*node.GetNode())
	bucket := routingTable.buckets[bucketIndex]
	bucket.AddToBucket(id, ip)
}

// getBucketIndex get the correct Bucket index for the KademliaID
func (routingTable *RoutingTable) getBucketIndex(node node.Node) int {
	distance := node.CalcDistance(routingTable.Me.ID)
	for i := 0; i < kademliaid.IDLength; i++ {
		for j := 0; j < 8; j++ {
			if (distance[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}

	return kademliaid.IDLength*8 - 1
}

// FindClosestContacts finds the count closest Contacts to the target in the RoutingTable
// **********************************************
// FIX THIS UP TO COMPLETE NODE LOOKUP
func (routingTable *RoutingTable) FindClosestContacts(target node.Node, count int) []node.Node {
	var candidates []node.Node
	bucketIndex := routingTable.getBucketIndex(target)
	bucket := routingTable.buckets[bucketIndex]

	candidates = append(candidates, bucket.GetContactAndCalcDistance(target)...)

	for i := 1; (bucketIndex-i >= 0 || bucketIndex+i < IDLength*8) && candidates.Len() < count; i++ {
		if bucketIndex-i >= 0 {
			bucket = routingTable.buckets[bucketIndex-i]
			candidates = append(candidates, bucket.GetContactAndCalcDistance(target)...)
		}
		if bucketIndex+i < IDLength*8 {
			bucket = routingTable.buckets[bucketIndex+i]
			candidates = append(candidates, bucket.GetContactAndCalcDistance(target)...)
		}
	}

	sort.Sort(candidates)

	if count > candidates.Len() {
		count = candidates.Len()
	}

	return candidates.GetContacts(count)
}
