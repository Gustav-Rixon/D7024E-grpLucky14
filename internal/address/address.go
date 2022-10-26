package address

import (
	"net"
	"os"

	"github.com/rs/zerolog/log"
)

type Address struct {
	host, port string
}

func New(address string) *Address {
	lport := os.Getenv("LISTEN_PORT")
	if lport == "" { // if the env var was not defined
		lport = "8888"
	}

	host, port, err := net.SplitHostPort(address)
	if err != nil && port == "" {
		parsedHost := net.ParseIP(address)
		if parsedHost == nil {
			log.Error().Msgf("Given address is not valid: %s", err)
			return &Address{}
		}
		host = parsedHost.String()
	}

	return &Address{
		host: host,
		port: lport,
	}
}

func (address *Address) String() string {
	return net.JoinHostPort(address.host, address.port)
}
