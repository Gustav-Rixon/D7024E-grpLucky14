package main

import (
	"fmt"
	"kademlia/internal/address"
	cmdlistener "kademlia/internal/command/listener"
	"kademlia/internal/constants"
	"kademlia/internal/network/listener"
	"kademlia/internal/network/restAPI"
	"kademlia/internal/node"
	"kademlia/internal/rpc/rpcqueue"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetHostIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal()
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
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
	logLevel := constants.LOG_LEVEL
	if err := InitLogger(logLevel); err == nil {
		log.Info().Str("Level", logLevel).Msg("Log level set")
	} else {
		log.Error().Str("Level", logLevel).Msg("Failed to parse log level, defaulting to info level...")
	}

	lport := constants.LISTEN_PORT

	host, err := os.Hostname()
	ip := GetHostIP().String()
	if err != nil {
		log.Error().Str("Host", host).Msgf("Failed to get container host: %s", err)
	}
	log.Info().Str("Hostname", host).Str("IP", ip).Msg("Starting node...")
	node := node.Node{}
	rpcQ := rpcqueue.New()
	addr := address.New(ip + ":8888")
	node.Init(addr)
	fmt.Println(node.ID)
	go cmdlistener.Listen(&node)
	go restAPI.Listen(&node)
	listener.Listen(ip, lport, &node, &rpcQ)

}
