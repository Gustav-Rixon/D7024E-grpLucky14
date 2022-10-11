package node

import (
	"fmt"
	"kademlia/internal/address"
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
	"strings"

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
	KClosest := node.RoutingTable.FindClosestContacts(key, requestorID, k)
	return KClosest
}

func (node *Node) NewRPC(content string, target *address.Address) rpc.RPC {
	return rpc.RPC{SenderId: node.ID, RPCId: kademliaid.NewRandomKademliaID(), Content: content, Target: target}
}

// https://kelseyc18.github.io/kademlia_vis/lookup/
func (node *Node) FIND_NODE(LookingUp *kademliaid.KademliaID) []contact.Contact {

	node.Shortlist = shortlist.NewShortlist(LookingUp, node.FindKClosest(node.ID, LookingUp, 5)) //INIT shortlist fullösning

	node.ProbeAlphaNodes(*node.Shortlist, 3)

	return node.Shortlist.GetContacts()

}

func (node *Node) FIND_DATA(hash *kademliaid.KademliaID) []contact.Contact {

	node.Shortlist = shortlist.NewShortlist(hash, node.FindKClosest(node.ID, hash, 5)) //INIT shortlist fullösning

	fmt.Println("SHORTLISTCREATED")
	fmt.Println(node.Shortlist)
	fmt.Println("SHORTLISTCREATED")

	fmt.Println("TARGET")
	fmt.Println(node.Shortlist.Target)
	fmt.Println("TARGET")

	node.ProbeAlphaNodesForData(*node.Shortlist, 3)

	return node.Shortlist.GetContacts()

}

func (node *Node) ProbeAlphaNodes(shortlist shortlist.Shortlist, alpha int) {

	numProbed := 0
	for i := 0; i < shortlist.Len() && numProbed < alpha; i++ {

		if !shortlist.Entries[i].Probed {
			log.Trace().Str("NodeID", shortlist.Entries[i].Contact.ID.String()).Msg("Probing node")
			shortlist.Entries[i].Probed = true
			rpc := node.NewRPC("FIND_NODE ", shortlist.Entries[i].Contact.Address)
			numProbed++
			node.Network.SendFindContactMessage(&rpc)
		}
	}
	node.RoutingTable.AddContact(*node.Shortlist.Closest)
}

func (node *Node) ProbeAlphaNodesForData(shortlist shortlist.Shortlist, alpha int) {

	numProbed := 0
	for i := 0; i < shortlist.Len() && numProbed < alpha; i++ {

		if !shortlist.Entries[i].Probed {
			log.Trace().Str("NodeID", shortlist.Target.String()).Msg("Probing node")
			shortlist.Entries[i].Probed = true
			rpc := node.NewRPC("FIND_VALUE ", shortlist.Entries[i].Contact.Address)
			numProbed++
			node.Network.SendFindDataMessage(&rpc)
		}
	}
	node.RoutingTable.AddContact(*node.Shortlist.Closest)
}
