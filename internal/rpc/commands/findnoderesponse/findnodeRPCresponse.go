package findnoderesponse

import (
	"errors"
	"fmt"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"strings"

	"github.com/rs/zerolog/log"
)

type FindNodeResponse struct {
	rpcId *kademliaid.KademliaID
	data  *string
}

func New(rpcId *kademliaid.KademliaID) *FindNodeResponse {
	return &FindNodeResponse{rpcId: rpcId}
}

// Will insert values to routing table
func (Resp *FindNodeResponse) Execute(node *node.Node) {
	log.Trace().Msg("Executing FIND_NODE_RESP RPC")

	//Just add contacts to its shortlist?
	response := *Resp.data
	desResponse := DeserializeContacts(response, node.ID)

	fmt.Println("@@@")
	fmt.Println(node.Shortlist.Entries)
	fmt.Println("@@@")

	//Insert all contacts to shortlist
	for _, element := range desResponse {
		node.Shortlist.Add(element)
	}

	fmt.Println("@@@2")
	//fmt.Println(node.Shortlist.Entries[0].Probed)
	//fmt.Println(node.Shortlist.Entries[1].Probed)
	fmt.Println("@@@2")

	node.ProbeAlphaNodes(*node.Shortlist, 3)

	//node.Shortlist = shortlist.NewShortlist(node.ID, desResponse)

	//desResponse := DeserializeContacts(response, node.Shortlist.Target)
	//Add the responses to the shorlist

	//Insert all contacts to shortlist
	//for _, element := range desResponse {
	//	node.Shortlist.Add(element)
	//}
	//fmt.Println("@@@")
	//fmt.Println(node.Shortlist)
	//fmt.Println(node.Shortlist.Entries)
	//fmt.Println("@@@")
}

func (Resp *FindNodeResponse) ParseOptions(options *[]string) error {
	if len(*options) == 0 {
		return errors.New("Empty RESP")
	}
	data := strings.Join(*options, " ")
	Resp.data = &data
	return nil
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
