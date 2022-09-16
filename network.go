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
	listenPort  string
	sendPort    string
}

// Borrowed code from Stack Overflow
// Gets the preferred outbound ip of this machine
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

// Gets initial network subnet info, and takes desired ports for i/o
func initNetwork(listenPort int, sendPort int) {
	var networkInfo NetworkInfo
	networkInfo.localIPAddr = getOutboundIP()
	networkInfo.listenPort = ":" + fmt.Sprint(listenPort) // Stored in format :port for ease of concatenating
	networkInfo.sendPort = ":" + fmt.Sprint(sendPort)     // Stored in format :port for ease of concatenating
	if listenPort < 0 || sendPort < 0 {
		log.Panicln("Invalid port number")
	}
	netInfo = networkInfo
}

// Basic function for eternally listening for Packets on a single port
func listen() {
	fmt.Println("Beginning to listen on ", netInfo.localIPAddr)
	for {
		message, senderAddr := awaitPacket()
		fmt.Println("Received message from ", senderAddr.IP.String(), "\n Sender ID: ", hex.EncodeToString(message.ID[:]))
		fmt.Println("Sending reply...")
		sendPacket(senderAddr.IP)
	}
}

// Function for pinging and awaiting reply from a given IP
// NOTE: Expects IP to exclude port number
func Ping(destIP net.IP) {
	fmt.Println("Pinging ", destIP)
	sendPacket(destIP)

	fmt.Println("Awaiting reply...")
	message, senderAddr := awaitPacket()
	fmt.Println("Received reply from ", senderAddr.String(), "\n Sender ID: ", hex.EncodeToString(message.ID[:]))
}

// Function for sending a (hardcoded) packet with current node's information to a given IP using UDP communication
func sendPacket(destIP net.IP) {

	// Begin by resolving the UDP address we should send from
	localAddr, err := net.ResolveUDPAddr("udp", netInfo.localIPAddr.String()+netInfo.sendPort)
	if err != nil {
		log.Fatal(err)
	}

	// Resolve destination UDP address
	destAddr, err := net.ResolveUDPAddr("udp", destIP.String()+netInfo.listenPort)
	if err != nil {
		log.Fatal(err)
	}

	// Establish UDP connection
	connection, err := net.DialUDP("udp", localAddr, destAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Packet and fill it with relevant info
	var sendPack Packet
	sendPack.ID = node.ID
	sendPack.IP = netInfo.localIPAddr

	// Encode Packet struct in order to send it across UDP connection
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encodeErr := encoder.Encode(sendPack)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	// Send Packet struct
	_, err = connection.Write(buffer.Bytes())
	if err != nil {
		fmt.Println("Ping failed: ", err)
	}

	connection.Close()
}

// Basic function for awaiting a single Packet, returns the Packet it gets along with the UDP address that sent it
//
// NOTE: MAKE SURE TO NEVER HAVE MORE THAN ONE AWAITPACKET ALIVE AT A TIME, waiting on the same port is highly illegal
// and will crash everything :(
func awaitPacket() (Packet, *net.UDPAddr) {

	// Resolve UDP address that we should listen on
	localAddress, err := net.ResolveUDPAddr("udp", netInfo.localIPAddr.String()+netInfo.listenPort)
	if err != nil {
		log.Fatal(err)
	}

	// Begin listening
	connection, err := net.ListenUDP("udp", localAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Declare Packet
	var message Packet

	// --- Could change everything below into some general handling of different types of packets ---

	// Create byte array "buffer" for storing incoming Packet
	inputBytes := make([]byte, 4096)
	length, senderAddr, _ := connection.ReadFromUDP(inputBytes)

	// Decode encoded Packet back into struct
	buffer := bytes.NewBuffer(inputBytes[:length])
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(&message) // <-- Stores the decoded Packet in message variable
	if err != nil {
		fmt.Println("Error decoding incoming packet: ", err)
	}

	connection.Close()
	return message, senderAddr
}

func handlePacket(p Packet) {
	switch p.pType {
	case ping:

		break
	}
}
