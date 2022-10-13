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

// Stores values in datastore, continually sends RefreshRPCs to other nodes with the same data on them as to not let it expire
func (node *Node) Store(value *string, contacts *[]contact.Contact) {
	log.Trace().Str("Value", *value).Msg("Storing value")
	node.DataStore.Insert(*value, contacts, node.Network.UdpSender)

	if contacts != nil {
		key := kademliaid.NewKademliaID(value)
		go func() {
			for {
				time.Sleep(datastore.TTL / 2)

				if node.DataStore.Refresh(key) == "" {
					break
				}

				for _, closeNode := range *contacts {
					rpc := node.NewRPC("REFRESH "+key.String(), closeNode.Address)
					node.Network.SendRefreshMessage(&rpc)
				}

			}
		}()
	}

}

func (node *Node) Refresh(key *string) {
	log.Trace().Str("Key", *key).Msg("Refreshing")
	//node.DataStore.Refresh(key)
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

	K, err := strconv.Atoi(os.Getenv("K"))
	if err != nil {
		log.Error().Msgf("Failed to convert env variable ALPHA from string to int: %s", err)
	}

	node.Shortlist = shortlist.NewShortlist(LookingUp, node.FindKClosest(node.ID, LookingUp, K)) //INIT shortlist fullösning

	for {
		closestSoFar := node.Shortlist.Closest
		probed := node.ProbeAlpha(*node.Shortlist, 4, fmt.Sprintf("%s %s", "FIND_NODE", LookingUp))

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

func (node *Node) FIND_DATA(hash *kademliaid.KademliaID) string {

	K, err := strconv.Atoi(os.Getenv("K"))
	if err != nil {
		log.Error().Msgf("Failed to convert env variable K from string to int: %s", err)
	}

	ALPHA, err := strconv.Atoi(os.Getenv("ALPHA"))
	if err != nil {
		log.Error().Msgf("Failed to convert env variable ALPHA from string to int: %s", err)
	}

	node.Shortlist = shortlist.NewShortlist(hash, node.FindKClosest(node.ID, hash, K)) //INIT shortlist fullösning

	result := ""

	for {
		closestSoFar := node.Shortlist.Closest
		probed := node.ProbeAlpha(*node.Shortlist, ALPHA, fmt.Sprintf("%s %s", fmt.Sprintf("FIND_VALUE %s", hash.String())))

		if probed == 0 {
			log.Trace().Msg("FIND_VALUE lookup became stale")
			break
		}

		if result != "" {
			return result
		}

		if node.Shortlist.Closest == closestSoFar {
			log.Trace().Msg("Closest node not updated")
			node.ProbeAlpha(*node.Shortlist, ALPHA, fmt.Sprintf("%s %s", fmt.Sprintf("FIND_VALUE %s", hash.String())))
			break
		}

	}

	s := "Value not found, k closest contacts: ["
	for i, entry := range node.Shortlist.Entries {
		if entry != nil {
			s += entry.Contact.String()
			if i < len(node.Shortlist.Entries)-1 {
				s += ", "
			}
		}
	}
	s += "]"
	return s

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
