// TODO pick: Switch cases or https://www.youtube.com/watch?v=7qKXxfs7LVY

package main

import (
	"kademlia/internal/kademliaid"
	"kademlia/internal/network"
	. "kademlia/internal/node"
	"kademlia/pkg/actions"
)

func main() {
	CreateSelf(kademliaid.NewRandomKademliaID(), network.GetOutboundIP())
	actions.Commands()
}
