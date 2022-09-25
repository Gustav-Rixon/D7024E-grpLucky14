package rpcparser

import (
	"errors"
	"fmt"
	"kademlia/internal/contact"
	"kademlia/internal/node"
	"kademlia/internal/rpc"
	"kademlia/internal/rpc/commands/ping"
	"strings"

	"github.com/rs/zerolog/log"
)

type RPCCommand interface {
	Execute(node *node.Node)
	ParseOptions(options *[]string) error
}

// Parses a rpc and returns a rpc command.
func ParseRPC(requestor *contact.Contact, rpc *rpc.RPC) (RPCCommand, error) {
	fields := strings.Fields(rpc.Content)
	if len(fields) == 0 {
		return nil, errors.New("Missing RPC name")
	}

	var cmd RPCCommand
	var err error
	rpcLog := log.Debug().Str("RPCId", rpc.RPCId.String())
	switch identifier := fields[0]; identifier {
	case "PING":
		rpcLog.Msg("PING received")
		cmd = ping.New(requestor.Address, rpc.RPCId)
	default:
		err = errors.New(fmt.Sprintf("Received unknown RPC %s", identifier))
		cmd = nil
	}
	return cmd, err
}
