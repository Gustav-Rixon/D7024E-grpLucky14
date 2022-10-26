package sender

import (
	"kademlia/internal/address"
	"kademlia/internal/constants"
	"net"
	"strconv"
)

type Sender struct {
	connection *net.UDPConn
}

func New() (*Sender, error) {
	senderPort := ":" + strconv.FormatInt(constants.SEND_PORT, 10) // Retrive port from docker env
	literalAddr, err := net.ResolveUDPAddr("", senderPort)
	connection, err := net.ListenUDP("udp4", literalAddr)
	if err != nil {
		return nil, err

	}
	return &Sender{connection: connection}, nil

}

func (udp *Sender) Send(data string, target *address.Address) error {
	address, err := net.ResolveUDPAddr("udp", target.String())
	if err != nil {
		return err
	}
	_, err = udp.connection.WriteTo([]byte(data), address)
	return err
}
