package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

type Node struct {
	ID *KademliaID
}

type Packet struct {
	ID [20]byte
	IP net.UDPAddr
}

func NewNode() Node {
	Id := NewRandomKademliaID()
	return Node{Id}
}

// Borrwed .)
// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// converts ip address to int
func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

// Returns random number, used in Kademlia ID generation
func getRandNum() int {
	r := rGen.Intn(256)
	return r
}

// Random number generator, use to get random numbers between nodes
var rGen *rand.Rand

// The node itself
var node Node

func main() {
	//initialize randomization of ID
	randSource := rand.NewSource(time.Now().UnixNano())
	rGen = rand.New(randSource)
	node.ID = NewRandomKademliaID()

	if GetOutboundIP().String() == "172.18.0.2" {
		listen()
	} else {
		send()
	}

	for {
		time.Sleep(2 * time.Second)

		fmt.Println(node.ID)

	}
}

func listen() {
	localAddress, err := net.ResolveUDPAddr("udp", GetOutboundIP().String()+":80")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Beginning to listen on ", localAddress)

	connection, err := net.ListenUDP("udp", localAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	var message Packet
	for {
		inputBytes := make([]byte, 4096)
		length, senderAddr, _ := connection.ReadFromUDP(inputBytes)

		buffer := bytes.NewBuffer(inputBytes[:length])
		decoder := gob.NewDecoder(buffer)
		err = decoder.Decode(&message)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Received message from ", senderAddr, "\n Packet IP: ", message.IP.String())
	}
}

func send() {
	dest_addr := "172.18.0.2"
	port := ":80"

	fmt.Printf("COMM: Broadcasting message to: %s%s\n", dest_addr, port)
	localAddr, err := net.ResolveUDPAddr("udp", GetOutboundIP().String()+port)

	if err != nil {
		log.Fatal(err)
	}

	sendAddr, err := net.ResolveUDPAddr("udp", dest_addr+port)

	if err != nil {
		log.Fatal(err)
	}

	connection, err := net.DialUDP("udp", localAddr, sendAddr)
	defer connection.Close()

	if err != nil {
		log.Fatal(connection, err)
	}

	//message := []byte(string("hello from " + node.ID.String()[0:4] + " :))))"))
	sendPack := Packet{}
	sendPack.ID = *node.ID
	sendPack.IP = *localAddr

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encodeErr := encoder.Encode(sendPack)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	for {
		fmt.Println("Sending packet")
		_, err = connection.Write(buffer.Bytes())
		if err != nil {
			fmt.Println("Write in broadcast localhost failed", err)
		}
		time.Sleep(2 * time.Second)
	}
}
