package listener

import (
	"kademlia/internal/address"
	"kademlia/internal/contact"
	"kademlia/internal/node"
	"kademlia/internal/rpc"
	"kademlia/internal/rpc/rpcparser"
	"kademlia/internal/rpc/rpcqueue"
	"net"
	"strings"

	"github.com/rs/zerolog/log"
)

//This package will handel the listening part. It will be writhen as an UDP.

// Sauce https://jameshfisher.com/2016/11/17/udp-in-go/ and https://pkg.go.dev/net
// Listen initiates a UDP server
func Listen(ip string, port int, node *node.Node, rpcQ *rpcqueue.RPCQueue) {
	addr := net.UDPAddr{IP: net.ParseIP(ip), Port: port}
	ln, err := net.ListenUDP("udp4", &addr)
	defer ln.Close()
	if err != nil {
		log.Error().Msgf("Failed to listen on UDP Address: %s", err)
	}
	log.Info().Str("Address", addr.String()).Msg("Listening on UDP packets on address")

	waitForMessages(ln, node, rpcQ)
}
func waitForMessages(con *net.UDPConn, node *node.Node, rpcQ *rpcqueue.RPCQueue) {
	for {
		//https://stackoverflow.com/questions/1098897/what-is-the-largest-safe-udp-packet-size-on-the-internet
		// I trust this post blindly
		udpBuffer := make([]byte, 512)
		nr, addr, err := con.ReadFromUDP(udpBuffer)
		if err != nil {
			log.Warn().Str("Error", err.Error()).Msg("Failed to read from UDP")
			return
		}
		//Go goroutine when message is read
		go func() { // Add goroutine when a rpc-message was read
			data := udpBuffer[0:nr]
			adr := address.New(addr.String())
			rpcMsg, err := rpc.Deserialize(string(data))
			if err == nil {
				c := contact.NewContact(rpcMsg.SenderId, adr)
				node.RoutingTable.AddContact(c)

				cmd, err := rpcparser.ParseRPC(&c, &rpcMsg)
				if err != nil {
					log.Warn().Str("Error", err.Error()).Msg("Failed to parse RPC")
				}

				options := strings.Split(rpcMsg.Content, " ")[1:]
				if err = cmd.ParseOptions(&options); err == nil {
					rpcQ.AddToQueue(node, cmd)
				} else {
					log.Warn().
						Str("Error", err.Error()).
						Msg("Failed to parse RPC options")
				}

				log.Trace().Str("NodeID", c.ID.String()).Str("Address", c.Address.String()).Msg("Inserting new node to bucket")
			} else {
				log.Warn().Str("Error", err.Error()).Msg("Failed to deserialize message in UDPListener")
			}
		}()
	}
}
