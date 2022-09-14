package bucket

import (
	"container/list"
	"fmt"
	. "kademlia/internal/Node"
	. "kademlia/internal/contact"
	. "kademlia/internal/kademliaid"
	. "kademlia/internal/network"
)

const bucketSize = 20

// bucket definition
// contains a List
type Bucket struct {
	list *list.List
}

// Creates a bucket
func newBucket() *Bucket {
	bucket := &Bucket{}
	bucket.list = list.New()
	return bucket
}

// Searches through a bucket for an ID from Packet, if the ID is found the entry corresponding to the ID get moved to the front of the buckets list
// OTHERWISE creates a new node instance and places it into the bucket
// THIS SHOULD BE CALLED in the listen function every time the node receives a message as per the kademlia specification.
func (b Bucket) addToBucket(p Packet) {
	var element *list.Element
	for e := b.list.Front(); e != nil; e = e.Next() {
		nodeID := e.Value.(Node).ID

		if (p).ID == nodeID {
			element = e
		}
	}
	if element == nil {
		if b.list.Len() < bucketSize {
			var n = NewNode(p.ID, p.IP)
			b.list.PushFront(n)
			fmt.Println("New Node added to bucket")
		}
	} else {
		b.list.MoveToFront(element)
		fmt.Println("Node found in bucket, moving to front")
	}
}

// *******OLD BUCKET CODE*********
type bucket2 struct {
	list *list.List
}

// newBucket returns a new instance of a bucket
func newBucket2() *bucket2 {
	bucket := &bucket2{}
	bucket.list = list.New()
	return bucket
}

// AddContact adds the Contact to the front of the bucket
// or moves it to the front of the bucket if it already existed
func (bucket *bucket2) AddContact(contact Contact) {
	var element *list.Element
	for e := bucket.list.Front(); e != nil; e = e.Next() {
		nodeID := e.Value.(Contact).ID

		if (contact).ID.Equals(nodeID) {
			element = e
		}
	}

	if element == nil {
		if bucket.list.Len() < bucketSize {
			bucket.list.PushFront(contact)
		}
	} else {
		bucket.list.MoveToFront(element)
	}
}

// GetContactAndCalcDistance returns an array of Contacts where
// the distance has already been calculated
func (bucket *bucket2) GetContactAndCalcDistance(target *KademliaID) []Contact {
	var contacts []Contact

	for elt := bucket.list.Front(); elt != nil; elt = elt.Next() {
		contact := elt.Value.(Contact)
		contact.CalcDistance(target)
		contacts = append(contacts, contact)
	}

	return contacts
}

// Len return the size of the bucket
func (bucket *bucket2) Len() int {
	return bucket.list.Len()
}
