package put

import (
	"errors"
	"fmt"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"strings"

	"github.com/rs/zerolog/log"
)

type Put struct {
	fileContent string
}

const k = 3

func (put *Put) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing put command")

	fmt.Println("!!!!!!!!!!!!!!!")
	fmt.Println(put.fileContent)
	fmt.Println("!!!!!!!!!!!!!!!")

	key := kademliaid.NewKademliaID(&put.fileContent)
	closestNodes := node.FindKClosest(&key, nil, k)
	node.Store(&put.fileContent, &closestNodes)

	for _, closeNode := range closestNodes {
		node.Network.SendStoreMessage(node.ID, closeNode.Address, []byte(put.fileContent))
	}

	return key.String(), nil
}

func (put *Put) ParseOptions(options []string) error {
	if len(options) < 1 {
		return errors.New("Missing file content")
	}
	put.fileContent = strings.Join(options[0:], " ")
	return nil
}

func (put *Put) PrintUsage() string {
	return "put <file content>"
}
