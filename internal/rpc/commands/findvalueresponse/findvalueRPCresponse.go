package findvalueresp

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
	awnser := strings.Split(Resp.content, "=")

	ALPHA := constants.ALPHA

	if awnser[0] == "VALUE" {
		log.Debug().Str("Key found", awnser[1])
		node.Shortlist.AddFoundData(Resp.senderId, awnser[1])

	} else {
		desResponse := DeserializeContacts(*response, node.ID)
		for _, element := range desResponse {
			node.Shortlist.Add(element)
		}

		node.ProbeAlpha(*node.Shortlist, ALPHA, fmt.Sprintf("%s %s", "FIND_NODE", node.Shortlist.Target))

		fmt.Println("@@@")
		fmt.Println(node.Shortlist)
		fmt.Println("@@@")
	}
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
