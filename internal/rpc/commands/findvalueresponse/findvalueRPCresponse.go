package findvalueresp

import (
	"errors"
	"fmt"
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"

	"strings"
)

type FindValueResp struct {
	content  string
	rpcId    *kademliaid.KademliaID
	senderId *kademliaid.KademliaID
}

func New(senderId *kademliaid.KademliaID, rpcId *kademliaid.KademliaID) *FindValueResp {
	return &FindValueResp{senderId: senderId, rpcId: rpcId}
}

func (Resp *FindValueResp) Execute(node *node.Node) {

	response := &Resp.content

	if *response == "" {
		desResponse := DeserializeContacts(*response, node.ID)
		for _, element := range desResponse {
			node.Shortlist.Add(element)
		}

		node.ProbeAlphaNodesForData(*node.Shortlist, 3)
	}

	fmt.Println("####")
	fmt.Println(*response)
	fmt.Println("####")
}

func (findresp *FindValueResp) ParseOptions(options *[]string) error {
	if len(*options) == 0 {
		return errors.New("missing content")
	}
	findresp.content = strings.Join((*options)[0:], " ")
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
