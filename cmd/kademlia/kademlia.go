package main

import (
	"kademlia/internal/address"
	cmdlistener "kademlia/internal/command/listener"
	"kademlia/internal/network/listener"
	"kademlia/internal/node"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func getHostIP() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		log.Error().Msgf("Failed to get container interface addresses: %s", err)
	}
	for _, address := range addresses {

		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func InitLogger(level string) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	zerolog.SetGlobalLevel(zerolog.InfoLevel) //Default to info level
	logLevel, err := zerolog.ParseLevel(level)
	if err == nil {
		zerolog.SetGlobalLevel(logLevel)
	}
	return err
}

func main() {
	logLevel := os.Getenv("LOG_LEVEL")
	if err := InitLogger(logLevel); err == nil {
		log.Info().Str("Level", logLevel).Msg("Log level set")
	} else {
		log.Error().Str("Level", logLevel).Msg("Failed to parse log level, defaulting to info level...")
	}

	lport, err := strconv.Atoi(os.Getenv("LISTEN_PORT"))
	if err != nil {
		log.Error().Msgf("Failed to convert env variable LISTEN_PORT from string to int: %s", err)
	}

	host, err := os.Hostname()
	ip := getHostIP()
	if err != nil {
		log.Error().Str("Host", host).Msgf("Failed to get container host: %s", err)
	}
	log.Info().Str("Hostname", host).Str("IP", ip).Msg("Starting node...")

	node := node.Node{}
	addr := address.New("localhost:8888")
	node.Init(addr)

	go cmdlistener.Listen(&node)
	listener.Listen(ip, lport, &node)
}
