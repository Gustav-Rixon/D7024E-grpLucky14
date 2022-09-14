package network

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

type Packet struct {
	ID [20]byte
	IP net.UDPAddr
}

type NetworkInfo struct {
	localIPAddr net.IP
	networkPort string
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
// Converts an ip address to int
func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

var netInfo NetworkInfo

// Gets initial network subnet info, and takes a desired port used for communication
func initNetwork(port int) {
	var networkInfo NetworkInfo
	networkInfo.localIPAddr = getOutboundIP()
	networkInfo.networkPort = ":" + fmt.Sprint(port) // Stored in format :port for ease of concatenating
	if port < 0 {
		log.Panicln("Invalid port number")
	}
	netInfo = networkInfo
}

func listen() {
	fmt.Println("Beginning to listen on:", netInfo.localIPAddr.String()+netInfo.networkPort)
	localAddress, err := net.ResolveUDPAddr("udp", netInfo.localIPAddr.String()+netInfo.networkPort)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Beginning to listen on ", localAddress)

	connection, err := net.ListenUDP("udp", localAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	var sendPack Packet
	sendPack.ID = node.ID
	sendPack.IP = *localAddress

	var replyBuffer bytes.Buffer
	encoder := gob.NewEncoder(&replyBuffer)
	encodeErr := encoder.Encode(sendPack)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	var message Packet
	for {
		// Could change this into some general handling of different types of packets
		inputBytes := make([]byte, 4096)
		length, senderAddr, _ := connection.ReadFromUDP(inputBytes)

		buffer := bytes.NewBuffer(inputBytes[:length])
		decoder := gob.NewDecoder(buffer)
		err = decoder.Decode(&message)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Received message from ", senderAddr, "\n Sender ID: ", hex.EncodeToString(message.ID[:]))
		fmt.Println("Sending reply...")
		_, err = connection.WriteToUDP(replyBuffer.Bytes(), senderAddr)
		if err != nil {
			fmt.Println("Ping failed: ", err)
		}
	}
}

// Function for pinging and awaiting reply from a given IP
// NOTE: Expects IP to exclude port number
func sendPing(destIP net.IP) {
	localAddr, err := net.ResolveUDPAddr("udp", netInfo.localIPAddr.String()+netInfo.networkPort)
	if err != nil {
		log.Fatal(err)
	}

	sendAddr, err := net.ResolveUDPAddr("udp", destIP.String()+netInfo.networkPort)
	if err != nil {
		log.Fatal(err)
	}

	connection, err := net.DialUDP("udp", localAddr, sendAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	var sendPack Packet
	sendPack.ID = node.ID
	sendPack.IP = *localAddr

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encodeErr := encoder.Encode(sendPack)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	fmt.Println("Pinging: ", destIP.String())
	_, err = connection.Write(buffer.Bytes())
	if err != nil {
		fmt.Println("Ping failed: ", err)
	}

	fmt.Println("Awaiting reply...")
	var message Packet

	inputBytes := make([]byte, 4096)
	length, senderAddr, _ := connection.ReadFromUDP(inputBytes)

	replyBuffer := bytes.NewBuffer(inputBytes[:length])
	decoder := gob.NewDecoder(replyBuffer)
	err = decoder.Decode(&message)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received reply from ", senderAddr, "\n Sender ID: ", hex.EncodeToString(message.ID[:]))
}

/*
func (info NetworkInfo) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (info NetworkInfo) SendFindDataMessage(hash string) {
	// TODO
}

func (info NetworkInfo) SendStoreMessage(data []byte) {
	// TODO
}
*/
