package findnoderesponse

import (
	"errors"
	"fmt"
	"kademlia/internal/constants"
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

	ALPHA := constants.ALPHA

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

	node.ProbeAlpha(*node.Shortlist, ALPHA, fmt.Sprintf("%s %s", "FIND_NODE", node.Shortlist.Target))

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
