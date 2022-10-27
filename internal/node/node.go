package node

import (
	"fmt"
	"kademlia/internal/address"
	"kademlia/internal/constants"
	"kademlia/internal/contact"
	"kademlia/internal/datastore"
	"kademlia/internal/kademliaid"
	"kademlia/internal/network"
	"kademlia/internal/network/sender"
	"kademlia/internal/routingtable"
	"kademlia/internal/rpc"
	"kademlia/internal/shortlist"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

// Used in main to call on NewRandomKademliaID function
type Node struct {
	ID           *kademliaid.KademliaID
	RoutingTable *routingtable.RoutingTable
	DataStore    datastore.DataStore
	Network      network.Network
	Shortlist    *shortlist.Shortlist //??????
	bootstrap    bool
}

// Init initializes the node by generating a NodeID and creating a
// data store
func (node *Node) Init(address *address.Address) {
	id := kademliaid.NewRandomKademliaID()
	me := contact.NewContact(id, address)
	Sender, err := sender.New()

	if err != nil {
		log.Fatal().Str("Error", err.Error()).Msg("Failed to initialize ndoe")
	}

	*node = Node{
		ID:           id,
		DataStore:    datastore.New(),
		RoutingTable: routingtable.NewRoutingTable(me),
		Network:      network.Network{UdpSender: Sender},
		Shortlist:    nil,
		bootstrap:    false,
	}
}

// Stores values in datastore, continually sends RefreshRPCs to other nodes with the same data on them as to not let it expire
func (node *Node) Store(value *string, contacts *[]contact.Contact) {
	log.Trace().Str("Value", *value).Msg("Storing value")
	node.DataStore.Insert(*value, contacts, node.Network.UdpSender, true)

	if contacts != nil {
		key := kademliaid.NewKademliaID(value)
		node.DataStore.SetForget(key, false)
		go func() {
			for {
				time.Sleep(datastore.TTL / 2)

				if node.DataStore.GetForget(key) {
					fmt.Println("Object will no longer be refreshed")
					break
				}

				node.DataStore.Refresh(key)

				for _, closeNode := range *contacts {
					rpc := node.NewRPC("REFRESH "+key.String(), closeNode.Address)
					node.Network.SendRefreshMessage(&rpc)
				}

			}
		}()
	}

}

// FindKClosest returns a list of candidates containing the k closest nodes
// to the key being searched for (from the nodes own bucket(s))
func (node *Node) FindKClosest(key *kademliaid.KademliaID, requestorID *kademliaid.KademliaID, k int) []contact.Contact {
	KClosest := node.RoutingTable.FindClosestContacts(key, requestorID, k)
	return KClosest
}

func (node *Node) NewRPC(content string, target *address.Address) rpc.RPC {
	return rpc.RPC{SenderId: node.ID, RPCId: kademliaid.NewRandomKademliaID(), Content: content, Target: target}
}

// https://kelseyc18.github.io/kademlia_vis/lookup/
func (node *Node) FIND_NODE(LookingUp *kademliaid.KademliaID) []contact.Contact {
	node.Shortlist = shortlist.NewShortlist(LookingUp, node.FindKClosest(node.ID, LookingUp, constants.K)) //INIT shortlist fullösning

	ALPHA, err := strconv.Atoi(os.Getenv("ALPHA"))
	if err != nil {
		log.Error().Msgf("Failed to convert env variable ALPHA from string to int: %s", err)
		ALPHA = 3
	}

	for {
		closestSoFar := node.Shortlist.Closest
		probed := node.ProbeAlpha(*node.Shortlist, ALPHA, fmt.Sprintf("%s %s", "FIND_NODE", LookingUp))

		if probed == 0 {
			log.Trace().Msg("FIND_NODE lookup became stale")
			break
		}

		if node.Shortlist.Closest == closestSoFar {
			log.Trace().Msg("Closest node not updated")
			node.ProbeAlpha(*node.Shortlist, 3, fmt.Sprintf("%s %s", "FIND_NODE", LookingUp))
			break
		}
	}

	fmt.Println("@@@@@@")
	fmt.Println(node.Shortlist)
	fmt.Println("@@@@@@")

	return node.Shortlist.GetContacts()

}

func (node *Node) FIND_DATA(hash *kademliaid.KademliaID) {

	ALPHA, err := strconv.Atoi(os.Getenv("ALPHA"))
	if err != nil {
		log.Error().Msgf("Failed to convert env variable ALPHA from string to int: %s", err)
		ALPHA = 3
	}

	node.Shortlist = shortlist.NewShortlist(hash, node.FindKClosest(node.ID, hash, constants.K)) //INIT shortlist fullösning

	for {
		closestSoFar := node.Shortlist.Closest
		probed := node.ProbeAlpha(*node.Shortlist, ALPHA, fmt.Sprintf("%s ", fmt.Sprintf("FIND_VALUE %s", hash.String())))

		if probed == 0 {
			log.Trace().Msg("FIND_VALUE lookup became stale")
			break
		}

		if node.Shortlist.Closest == closestSoFar {
			log.Trace().Msg("Closest node not updated")
			node.ProbeAlpha(*node.Shortlist, ALPHA, fmt.Sprintf("%s ", fmt.Sprintf("FIND_VALUE %s", hash.String())))
			break
		}

	}

}

// ROUND X
func (node *Node) ProbeAlpha(shortlist shortlist.Shortlist, alpha int, content string) int {

	numProbed := 0
	for i := 0; i < shortlist.Len() && numProbed < alpha; i++ {

		if !shortlist.Entries[i].Probed {
			log.Trace().Str("NodeID", shortlist.Entries[i].Contact.ID.String()).Msg("Probing node")
			shortlist.Entries[i].Probed = true
			rpc := node.NewRPC(content, shortlist.Entries[i].Contact.Address)
			numProbed++
			node.Network.SendFindContactMessage(&rpc)
		}
	}

	return numProbed
}
