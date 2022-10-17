package get

//TODO GET THE BREaD
import (
	"errors"
	"fmt"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"time"

	"github.com/rs/zerolog/log"
)

type Get struct {
	hash kademliaid.KademliaID
}

func (get *Get) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing get command")
	// Check local storage
	value := node.DataStore.GetValue(get.hash)
	if value != "" {
		value += ", from local node"
	} else {
		log.Debug().Str("Key", get.hash.String()).Msg("Value not found locally")
		node.FIND_DATA(&get.hash)
		timeStamp := time.Now()

		for node.Shortlist.GetData() == "" {
			time.Sleep(1 * time.Millisecond)
			if time.Since(timeStamp) > (5 * time.Second) { //WORST CASE
				break
			}
		}

		//fmt.Println("@@@@@@@test@@@@@@@")
		//fmt.Println(t)
		//fmt.Println("@@@@@@@test@@@@@@@")

		//value = node.Shortlist.GetData() + " from "
		value = node.Shortlist.GetData()

		fmt.Println("@@@@@CLOSETS TARGET INFO@@@@@")
		fmt.Println(node.Shortlist.Entries[0].Contact.Address)
		fmt.Println(node.Shortlist.Entries[0].Probed)

		fmt.Println("@@@@@CLOSETS TARGET INFO@@@@@")

	}
	if value == "" {
		node.FIND_DATA(&get.hash)
		timeStamp := time.Now()

		for node.Shortlist.GetData() == "" {
			time.Sleep(1 * time.Millisecond)
			if time.Since(timeStamp) > (5 * time.Second) { //WORST CASE
				break
			}
		}

		value = node.Shortlist.GetData()

		if value != "" {
			return value, nil
		} else {
			return "", errors.New("Key not found")
		}

	}
	return value, nil
}

func (get *Get) ParseOptions(options []string) error {
	if len(options) < 1 {
		return errors.New("Missing hash")
	}
	get.hash = *kademliaid.FromString(options[0])
	return nil
}

func (get *Get) PrintUsage() string {
	return "USAGE: get <hash>"
}
