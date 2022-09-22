package routingtable

import (
	"kademlia/internal/bucket"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"net"
)

const bucketSize = 20

// Routing table used for testing
var rt RoutingTable

// Definitely move this
func GetRT() *RoutingTable {
	return &rt
}

// RoutingTable definition
// keeps a refrence contact of me and an array of buckets
type RoutingTable struct {
	Me      node.Node
	buckets [kademliaid.IDLength * 8]*bucket.Bucket
}

// NewRoutingTable returns a new instance of a RoutingTable
func NewRoutingTable(me node.Node) *RoutingTable {
	routingTable := &RoutingTable{}
	for i := 0; i < kademliaid.IDLength*8; i++ {
		routingTable.buckets[i] = bucket.NewBucket()
	}
	routingTable.Me = me
	return routingTable
}

// Creates a new node instance, used when adding a node to bucket
// to transform the info from the message into a node instance
func NewNode(id [kademliaid.IDLength]byte, ip net.IP) node.Node {
	Id := kademliaid.NewKademliaID(id)
	//fmt.Println("Successfully created instance of Kademlia ID: ", *Id, " With IP: ", ip.String())
	return node.Node{Id, ip}
}

// AddContact add a new contact to the correct Bucket
func (routingTable *RoutingTable) AddContact(id [kademliaid.IDLength]byte, ip net.IP) {
	bucketIndex := routingTable.getBucketIndex(*node.GetNode())
	bucket := routingTable.buckets[bucketIndex]
	bucket.AddToBucket(id, ip)
}

// getBucketIndex get the correct Bucket index for the KademliaID
func (routingTable *RoutingTable) getBucketIndex(node node.Node) int {
	distance := node.CalcDistance(rt.Me.ID)
	for i := 0; i < kademliaid.IDLength; i++ {
		for j := 0; j < 8; j++ {
			if (distance[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}

	return kademliaid.IDLength*8 - 1
}

/*

// FindClosestContacts finds the count closest Contacts to the target in the RoutingTable
func (routingTable *RoutingTable) FindClosestContacts(target *KademliaID, count int) []Contact {
	var candidates ContactCandidates
	bucketIndex := routingTable.getBucketIndex(target)
	bucket := routingTable.buckets[bucketIndex]

	candidates.Append(bucket.GetContactAndCalcDistance(target))

	for i := 1; (bucketIndex-i >= 0 || bucketIndex+i < IDLength*8) && candidates.Len() < count; i++ {
		if bucketIndex-i >= 0 {
			bucket = routingTable.buckets[bucketIndex-i]
			candidates.Append(bucket.GetContactAndCalcDistance(target))
		}
		if bucketIndex+i < IDLength*8 {
			bucket = routingTable.buckets[bucketIndex+i]
			candidates.Append(bucket.GetContactAndCalcDistance(target))
		}
	}

	candidates.Sort()

	if count > candidates.Len() {
		count = candidates.Len()
	}

	return candidates.GetContacts(count)
}

*/
