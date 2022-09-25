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
	}
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

func (node *Node) FindData(hash *kademliaid.KademliaID) string {
	//TODO CREATE NODELOOKUP
	return "ajajajaj"

}

// Node lookup retunrs the k closest node's to the given key
// WIll HARD CODE THE SHIT OUT OF THIS - GURX
// What does the lookup need to have?
//
//	Probe K NODES FUCK MY LIFE
func (node *Node) nodeLookup(id *kademliaid.KademliaID) []contact.Contact {

	for {

	}
}

func (node *Node) probeNode() {

}
