package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

type NetworkInfo struct {
	localIPAddr net.IP
	sendPort    string
	listenPort  string
	outUDP      net.UDPAddr // Resolved UDP address to send Packets from
	inUDP       net.UDPAddr // Resolved UDP address used to listen for Packets with
}

// Borrowed code from Stack Overflow
// Gets the preferred outbound ip of this machine/container
func getOutboundIP() net.IP {
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
var netInfo NetworkInfo

// Initializes global variable netInfo
func initNetwork(listenPort int, sendPort int) {
	var networkInfo NetworkInfo

	// Get basic info

	networkInfo.localIPAddr = getOutboundIP()
	networkInfo.sendPort = ":" + fmt.Sprint(sendPort)     // Stored in format :port for ease of concatenating
	networkInfo.listenPort = ":" + fmt.Sprint(listenPort) // ditto
	if listenPort < 0 || sendPort < 0 {
		log.Panicln("Invalid port number")
	}

	// Resolve local UDP addresses

	outUDP, err := net.ResolveUDPAddr("udp", networkInfo.localIPAddr.String()+networkInfo.sendPort)
	if err != nil {
		log.Fatal(err)
	}

	inUDP, err := net.ResolveUDPAddr("udp", networkInfo.localIPAddr.String()+networkInfo.listenPort)
	if err != nil {
		log.Fatal(err)
	}

	networkInfo.outUDP = *outUDP
	networkInfo.inUDP = *inUDP

	netInfo = networkInfo
}

// Function that eternally listens for- and handles Packets
func listen() {
	fmt.Println("Beginning to listen on ", netInfo.localIPAddr)

	var p Packet
	for {
		p = awaitPacket()
		go handlePacket(p) // Handle the packet on a separate thread to avoid wasting time/missing packets
	}
}

// ====================================================
//
// REMEMBER TO REMOVE AWAITPACKET CALL (SUCKS!!!!!!!!!)
//
// ====================================================
//
// Function for pinging and awaiting reply from a given IP
// NOTE: Expects IP to exclude port number
func sendPing(destIP net.IP) {
	fmt.Println("Pinging ", destIP)
	sendPacket(createPingPacket(node), destIP)

	fmt.Println("Awaiting reply...")
	p := awaitPacket()
	fmt.Println("receieved replyey: ", p.pType)
	handlePacket(p)
}

// Function for sending a given packet 'p' to a given IP using UDP communication
func sendPacket(p Packet, destIP net.IP) {
	// Resolve destination UDP address
	destAddr, err := net.ResolveUDPAddr("udp", destIP.String()+netInfo.listenPort)
	if err != nil {
		log.Fatal(err)
	}

	// Establish UDP connection
	connection, err := net.DialUDP("udp", &netInfo.outUDP, destAddr)
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
	connection, err := net.ListenUDP("udp", &netInfo.inUDP)
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
	switch p.pType {
	case ping:
		fmt.Println("Received ping from ", p.IP.String(), "\n Sender ID: ", hex.EncodeToString(p.ID[:]))

		nodes := []Node{NewNode(NewRandomKademliaID(), p.IP),
			NewNode(NewRandomKademliaID(), p.IP),
			NewNode(NewRandomKademliaID(), p.IP),
			NewNode(NewRandomKademliaID(), p.IP)}

		sendPacket(createReturnNodesPacket(node, nodes), p.IP)
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
		fmt.Println("reading nodes")
		for _, n := range p.nodes {
			fmt.Println(n)
		}
		// TODO
		break

	case return_value:
		// TODO
		break

	default:
		log.Fatal("Unknown package type received")
	}
}
