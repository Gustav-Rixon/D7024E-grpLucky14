package main

import (
	"kademlia/pkg/actions"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

func main() {

	if len(os.Args) > 1 { // If a command was specified
		msg := StrArrayToByteArray(os.Args[1:])
		actions.Commands(&msg)
	} else {
		//TODO: Print usage
		log.Print("Usage: To be done...")
	}
}

func StrArrayToByteArray(strs []string) []byte {
	return []byte(strings.Join(strs, " "))
}
