package cmdparser

import (
	"strings"

	"kademlia/internal/command/commands/add"
<<<<<<< Updated upstream
	"kademlia/internal/command/commands/findnode"
=======
	"kademlia/internal/command/commands/forget"
>>>>>>> Stashed changes
	"kademlia/internal/command/commands/get"
	"kademlia/internal/command/commands/getTable"
	"kademlia/internal/command/commands/getid"
	"kademlia/internal/command/commands/join"
	"kademlia/internal/command/commands/ping"
	"kademlia/internal/command/commands/put"

	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

type Command interface {
	Execute(node *node.Node) (string, error)

	// Parse the options (i.e. words after command) and set related fields in
	// the struct
	ParseOptions(options []string) error

	PrintUsage() string
}

func ParseCmd(s string) Command {
	fields := strings.Fields(s)

	var command Command

	// Assume the string has already been checked to contain a command
	switch cmd := fields[0]; cmd {
	case "ping":
		command = new(ping.Ping)

	case "put":
		command = new(put.Put)

	case "get":
		command = new(get.Get)

	case "add":
		command = new(add.Add)

	case "getTable":
		command = new(getTable.GetTable)

	case "getid":
		command = new(getid.GetId)
	case "join":
		command = new(join.Join)

<<<<<<< Updated upstream
	//case "findValue":
	//	command = new()
=======
	case "forget":
		command = new(forget.Forget)
>>>>>>> Stashed changes

	default:
		log.Error().Str("command", cmd).Msg("Received unknown command")
		return nil
	}

	err := command.ParseOptions(fields[1:]) // Extract all options
	if err != nil {
		log.Error().Str("error", err.Error()).Msg("Failed to parse options")
		log.Info().Msg(command.PrintUsage())
		return nil
	}

	return command
}
