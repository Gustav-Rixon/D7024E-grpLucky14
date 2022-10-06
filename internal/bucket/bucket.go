package bucket

import (
	"container/list"
	. "kademlia/internal/contact"
	. "kademlia/internal/kademliaid"
)

// bucket definition
// contains a List
type Bucket struct {
	list *list.List
}

const bucketSize = 20

// NewBucket returns a new instance of a Bucket
func NewBucket() *Bucket {
	bucket := &Bucket{}
	bucket.list = list.New()
	return bucket
}

// Searches through a bucket for an ID from Packet, if the ID is found the entry corresponding to the ID get moved to the front of the buckets list
// OTHERWISE creates a new node instance and places it into the bucket
// THIS SHOULD BE CALLED in the listen function every time the node receives a message as per the kademlia specification.
func (b Bucket) AddToBucket(id [IDLength]byte, ip net.IP) {
	var element *list.Element
	for e := b.list.Front(); e != nil; e = e.Next() {
		nodeID := e.Value.(Node).ID

		if id == nodeID {
			element = e
		}
	}
	if element == nil {
		if b.list.Len() < bucketSize {
			n := NewNode(id, ip)
			b.list.PushFront(n)
			fmt.Println("New Node added to bucket")
			n.CalcDistance(node.GetNode().ID)
			fmt.Println("Distance from me to node = ", n.GetDistance())
		}
	} else {
		b.list.MoveToFront(element)
		fmt.Println("Node found in bucket, moving to front")
	}
}

// GetContactAndCalcDistance returns an array of Contacts where
// the distance has already been calculated
func (bucket *Bucket) GetContactAndCalcDistance(target node.Node) []node.Node {
	var contacts []node.Node

	for elt := bucket.list.Front(); elt != nil; elt = elt.Next() {
		node := elt.Value.(Node)
		node.CalcDistance(target.ID)
		contacts = append(contacts, node)
	}

	return contacts
}


// AddContact adds the Contact to the front of the bucket
// or moves it to the front of the bucket if it already existed
func (bucket *Bucket) AddContact(contact Contact) {
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
func (bucket *Bucket) GetContactAndCalcDistance(target *KademliaID) []Contact {
	var contacts []Contact

	for elt := bucket.list.Front(); elt != nil; elt = elt.Next() {
		contact := elt.Value.(Contact)
		contact.CalcDistance(target)
		contacts = append(contacts, contact)
	}

	return contacts
}

// GetContactAndCalcDistance returns an array of Contacts where the distance
// has already been calculated. This array will never contain a contact with
// the same nodeID as the requestorID.
func (bucket *Bucket) GetContactAndCalcDistanceNoRequestor(target *KademliaID, requestorID *KademliaID) []Contact {
	var contacts []Contact

	for elt := bucket.list.Front(); elt != nil; elt = elt.Next() {
		contact := elt.Value.(Contact)
		if !contact.ID.Equals(requestorID) {
			contact.CalcDistance(target)
			contacts = append(contacts, contact)
		}
	}
	return contacts
}

// Len return the size of the bucket
func (bucket *Bucket) Len() int {
	return bucket.list.Len()
}
