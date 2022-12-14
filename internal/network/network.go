package network

import (
	"fmt"
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"kademlia/internal/network/sender"
	"kademlia/internal/rpc"

	"github.com/rs/zerolog/log"
)

type Network struct {
	UdpSender *sender.Sender
}

// https://gist.github.com/night-codes/9a0f2a1933eb06a0feea9f940c09d3d4
// SendPongMessage replies a "PONG" message to the remote "pinger" address
func (network *Network) SendPongMessage(senderId *kademliaid.KademliaID, target *address.Address, id *kademliaid.KademliaID) error {
	rpc := rpc.New(senderId, "PONG", target)
	rpc.RPCId = id
	err := rpc.Send(network.UdpSender, target)

	if err != nil {
		log.Error().Msgf("Failed to write RPC PING message to UDP: %s", err.Error())
	}
	log.Debug().Str("Target", target.String()).Msg("Sent PONG RPC to target")

	return err
}

// SendPingMessage sends a "PING" message to a remote address
func (network *Network) SendPingMessage(senderId *kademliaid.KademliaID, target *address.Address) error {
	rpc := rpc.New(senderId, "PING", target)

	err := rpc.Send(network.UdpSender, target)
	if err != nil {
		log.Error().Msgf("Failed to write RPC PING message to UDP: %s", err.Error())
	}
	log.Debug().Str("Target", target.String()).Msg("Sent PING to target")

	return err
}

func (network *Network) SendFindContactMessage(rpc *rpc.RPC) error {
	err := rpc.Send(network.UdpSender, rpc.Target)
	if err != nil {
		log.Error().Msgf("Failed to write FIND_NODE RPC to UDP: %s", err.Error())
	}
	log.Debug().Str("Target", rpc.Target.String()).Str("rpcId", rpc.RPCId.String()).Msg("Sent FIND_NODE RPC to target")

	return err
}

// SendFindContactRespMessage responds to a FIND_NODE RPC by returning the k
// closest contacts to the key that the node knows of
func (network *Network) SendFindContactRespMessage(senderId *kademliaid.KademliaID, target *address.Address, rpcId *kademliaid.KademliaID, content *string) error {

	rpc := rpc.NewWithID(senderId, fmt.Sprintf("%s %s", "FIND_NODE_RESP", *content), target, rpcId)

	err := rpc.Send(network.UdpSender, target)
	if err != nil {
		log.Error().Msgf("Failed to write FIND_NODE_RESP message to UDP: %s", err.Error())
	}
	log.Debug().Str("Target", target.String()).Msg("FIND_NODE_RESP sent to target")

	return err
}

func (network *Network) SendFindDataMessage(rpc *rpc.RPC) error {
	err := rpc.Send(network.UdpSender, rpc.Target)

	if err != nil {
		log.Error().Msgf("Failed to write RPC FIND_VALUE message to UDP: %s", err.Error())
	}
	log.Debug().Str("Target", rpc.Target.String()).Msg("Sent FIND_VALUE RPC to target")

	return err
}

func (network *Network) SendFindDataRespMessage(senderID *kademliaid.KademliaID, target *address.Address, rpcId *kademliaid.KademliaID, content *string) error {
	rpc := rpc.NewWithID(senderID, fmt.Sprintf("FIND_VALUE_RESP %s", *content), target, rpcId)
	err := rpc.Send(network.UdpSender, target)

	if err != nil {
		log.Error().Msgf("Failed to write RPC FIND_VALUE_RESP message to UDP: %s", err.Error())
	}
	log.Debug().Str("Target", target.String()).Msg("Sent FIND_VALUE_RESP RPC to target")

	return err
}

func (network *Network) SendStoreMessage(senderId *kademliaid.KademliaID, target *address.Address, data []byte) error {
	rpc := rpc.New(senderId, fmt.Sprintf("%s %s", "STORE", data), target)
	err := rpc.Send(network.UdpSender, target)

	if err != nil {
		log.Error().Msgf("Failed to write RPC STORE message to UDP: %s", err.Error())
	}
	log.Debug().Str("Target", target.String()).Msg("Sent STORE RPC to target")

	return err
}

func (network *Network) SendRefreshMessage(rpc *rpc.RPC) error {
	err := rpc.Send(network.UdpSender, rpc.Target)

	if err != nil {
		log.Error().Msgf("Failed to write RPC REFRESH message to UDP: %s", err.Error())
		return err
	}
	log.Debug().Str("Target", rpc.Target.String()).Msg("Sent REFRESH RPC to target")

	return err
}
