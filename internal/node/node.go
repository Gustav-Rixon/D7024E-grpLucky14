package node

import (
	"kademlia/internal/address"
	"kademlia/internal/contact"
	"kademlia/internal/datastore"
	"kademlia/internal/kademliaid"
	"kademlia/internal/network"
	"kademlia/internal/network/sender"
	"kademlia/internal/routingtable"
	"kademlia/internal/rpc"
	"kademlia/internal/rpcqueue"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// Used in main to call on NewRandomKademliaID function
type Node struct {
	ID           *kademliaid.KademliaID
	RPCQueue     *rpcqueue.RPCQueue
	RoutingTable *routingtable.RoutingTable
	DataStore    datastore.DataStore
	Network      network.Network
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
		RPCQueue:     rpcqueue.New(),
		Network:      network.Network{UdpSender: Sender},
		bootstrap:    false,
	}

}

// SUPER/BOOT will have the ID: 9514e18b679622b8d59991a6298559cb03099d64
func (node *Node) InitBOOT(address *address.Address) {
	skitt := "0000000000000000000000000000000000000000"
	id := kademliaid.NewKademliaID(&skitt)
	me := contact.NewContact(&id, address)
	Sender, err := sender.New()

	if err != nil {
		log.Fatal().Str("Error", err.Error()).Msg("Failed to initialize ndoe")
	}

	*node = Node{
		ID:           &id,
		DataStore:    datastore.New(),
		RoutingTable: routingtable.NewRoutingTable(me),
		RPCQueue:     rpcqueue.New(),
		Network:      network.Network{UdpSender: Sender},
		bootstrap:    true,
	}
}

func (node *Node) AddRout(address *address.Address) {
	skitt := "0000000000000000000000000000000000000000"
	id := kademliaid.NewKademliaID(&skitt)
	me := contact.NewContact(&id, address)
	node.RoutingTable.AddContact(me)
}

func (node *Node) Store(value *string, contacts *[]contact.Contact) {
	log.Trace().Str("Value", *value).Msg("Storing value")
	node.DataStore.Insert(*value, contacts, node.Network.UdpSender)
}

func GetEnvIntVariable(variable string, defaultValue int) int {
	val, err := strconv.Atoi(os.Getenv(variable))
	if err != nil {
		log.Error().Msgf("Failed to convert env variable %s from string to int: %s", variable, err)
		return defaultValue
	}
	return val
}

func DeserializeContacts(data string, targetId *kademliaid.KademliaID) []*contact.Contact {
	contacts := []*contact.Contact{}
	for _, sContact := range strings.Split(data, " ") {
		if sContact != "" {
			err, c := contact.Deserialize(&sContact)
			if err == nil {
				c.CalcDistance(targetId)
				contacts = append(contacts, c)
			}
		}
	}
	return contacts
}

// FindKClosest returns a list of candidates containing the k closest nodes
// to the key being searched for (from the nodes own bucket(s))
func (node *Node) FindKClosest(key *kademliaid.KademliaID, requestorID *kademliaid.KademliaID, k int) []contact.Contact {
	return node.RoutingTable.FindClosestContacts(key, requestorID, k)
}

func (node *Node) NewRPC(content string, target *address.Address) rpc.RPC {
	return rpc.RPC{SenderId: node.ID, RPCId: kademliaid.NewRandomKademliaID(), Content: content, Target: target}
}

func (node *Node) NodeLookup(hash *kademliaid.KademliaID) []contact.Contact {
	//TODO CREATE NODELOOKUP SO THAT It CAN find close NODES THAT ARE NOT IN TABLE
	candidats := node.FindKClosest(hash, node.ID, 3)

	/*

		for each node Send an RPC telling them to do a nodeLookup to
			Each routine will check its address for TARGET(node,value,idk) locally
				If TARGET found GG. Send an RPC call
				Else Forward to its ALPHA NODES



	*/

	return candidats

}

// Finds data
/*
func (node *Node) FindData(hash *kademliaid.KademliaID) string {

	data := ""
	for {

	}
}
*/
