package network

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"kademlia/internal/node"
	"log"
	"net"
)

type NetworkInfo struct {
	LocalIPAddr net.IP
	sendPort    string
	listenPort  string
	outUDP      net.UDPAddr // Resolved UDP address to send Packets from
	inUDP       net.UDPAddr // Resolved UDP address used to listen for Packets with
}

// Borrowed code from Stack Overflow
// Gets the preferred outbound ip of this machine/container
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// Also borrowed from Stack Overflow
// Converts an ip address to int (useful??????)
/* func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
} */

// Global variable holding all useful container network information for communication between containers
var NetInfo NetworkInfo

// Initializes global variable netInfo
func InitNetwork(listenPort int, sendPort int) error {
	var networkInfo NetworkInfo

	// Get basic info

	networkInfo.LocalIPAddr = GetOutboundIP()
	networkInfo.sendPort = ":" + fmt.Sprint(sendPort)     // Stored in format :port for ease of concatenating
	networkInfo.listenPort = ":" + fmt.Sprint(listenPort) // ditto

	// Resolve local UDP addresses

	outUDP, err := net.ResolveUDPAddr("udp", networkInfo.LocalIPAddr.String()+networkInfo.sendPort)
	if err != nil {
		fmt.Println("Bad send address during network init")
		return err
	}

	inUDP, err := net.ResolveUDPAddr("udp", networkInfo.LocalIPAddr.String()+networkInfo.listenPort)
	if err != nil {
		fmt.Println("Bad listen address during network init")
		return err
	}

	networkInfo.outUDP = *outUDP
	networkInfo.inUDP = *inUDP

	NetInfo = networkInfo
	return nil
}

// Function that eternally listens for- and handles Packets
func Listen() {
	fmt.Println("Beginning to listen on ", NetInfo.LocalIPAddr)

	var p Packet
	for {
		p = awaitPacket()
		go handlePacket(p) // Handle the packet on a separate thread to avoid wasting time/missing packets
	}
}

// Function for pinging a given IP
//
// TODO:
// Once routing is implemented, change to ping a specific Node
func SendPing(destIP net.IP) {
	fmt.Println("Pinging ", destIP)
	sendPacket(createPingPacket(*node.GetNode()), destIP)
}

// Function for sending a given packet 'p' to a given IP using UDP communication
func sendPacket(p Packet, destIP net.IP) {
	// Resolve destination UDP address
	destAddr, err := net.ResolveUDPAddr("udp", destIP.String()+NetInfo.listenPort)
	if err != nil {
		log.Fatal(err)
	}

	// Establish UDP connection
	connection, err := net.DialUDP("udp", &NetInfo.outUDP, destAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Encode p in order to send it across UDP connection
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encodeErr := encoder.Encode(p)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	// Send p
	_, err = connection.Write(buffer.Bytes())
	if err != nil {
		fmt.Println("Could not send packet: ", err)
	}

	connection.Close()
}

// Basic function for awaiting a single Packet. Blocks thread until a Packet is received and then returns it.
//
// NOTE: MAKE SURE TO NEVER HAVE MORE THAN ONE AWAITPACKET ALIVE AT A TIME, waiting on the same port is highly illegal
// and will crash everything :(
func awaitPacket() Packet {
	// Begin listening
	connection, err := net.ListenUDP("udp", &NetInfo.inUDP)
	if err != nil {
		log.Fatal(err)
	}

	var p Packet

	// Create byte array "buffer" for storing incoming Packet
	inputBytes := make([]byte, 4096)
	length, _, _ := connection.ReadFromUDP(inputBytes)

	// Decode encoded Packet back into struct
	buffer := bytes.NewBuffer(inputBytes[:length])
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(&p) // <-- Stores the decoded Packet in p
	if err != nil {
		fmt.Println("Error decoding incoming packet: ", err)
	}

	connection.Close()
	return p
}

func handlePacket(p Packet) {
	switch p.PType {
	case ping:
		fmt.Println("Received ping from ", p.IP.String(), "\nSender ID: ", hex.EncodeToString(p.ID[:]))
		sendPacket(createACKPacket(*node.GetNode()), p.IP)
		break

	case find_node:
		// TODO, probably some routing inside the routing package, followed by sending a return_nodes Packet
		break

	case find_value:
		// TODO
		break

	case store:
		// TODO
		break

	case ACK:
		fmt.Println("Received reply from ", p.IP.String(), "\nSender ID: ", hex.EncodeToString(p.ID[:]))
		break

	case return_nodes:
		// TODO
		break

	case return_value:
		// TODO
		break

	default:
		log.Fatal("Unknown package type received: ", p.PType)
	}
}
