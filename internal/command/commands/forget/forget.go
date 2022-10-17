package forget

import (
	"errors"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"strings"

	"github.com/rs/zerolog/log"
)

type Forget struct {
	fileContent string
}

func (forg *Forget) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing forget command")
	key := kademliaid.FromString(forg.fileContent)
	node.DataStore.SetForget(*key, true)

	return key.String(), nil
}

func (forg *Forget) ParseOptions(options []string) error {
	if len(options) < 1 {
		return errors.New("Missing file content")
	}
	forg.fileContent = strings.Join(options[0:], " ")
	return nil
}

func (forg *Forget) PrintUsage() string {
	return "forget <file content>"
}
