package main

import (
	"math/rand"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Used in main to call on NewRandomKademliaID function
type Node struct {
	ID [IDLength]byte
	IP net.UDPAddr
}

func NewNode() Node {
	Id := NewRandomKademliaID()
	return Node{Id}
=======
type Packet struct {
	ID [IDLength]byte
	IP net.UDPAddr
}

func NewNode(id [IDLength]byte, ip net.UDPAddr) Node {
	Id := NewKademliaID(id)
	//fmt.Println("Successfully created instance of Kademlia ID: ", *Id, " With IP: ", ip.String())
	return Node{Id, ip}
}

// Returns random number, used in Kademlia ID generation
func getRandNum() int {
	r := rGen.Intn(256)
	return r
}

// Takes two binary numbers and does a XOR b
func getDistance(a int, b int) int {
	distance := a ^ b
	return distance
}

// Random number generator, use to get random numbers between nodes
var rGen *rand.Rand

// The node itself
var node Node

// Bucket used for testing
var b *Bucket

func main() {
	// initialize randomization of ID
	randSource := rand.NewSource(time.Now().UnixNano())
	rGen = rand.New(randSource)
	node.ID = NewRandomKademliaID()

	xt := reflect.TypeOf(node.ID).Kind()
	xtx := reflect.TypeOf(*node.ID).Kind()
	dist := getDistance(00, 10)

	for {
		fmt.Print("node Id:")
		fmt.Println(node.ID)
		fmt.Print("node Id type:")
		fmt.Println(xt)
		fmt.Println(*node.ID)
		fmt.Println(xt)
		fmt.Println(xtx)
		fmt.Println(dist)
		time.Sleep(1 * time.Second)
	}
}

	// initialize network settings, communicate via port 80
	initNetwork(80)

	if netInfo.localIPAddr.Mask(net.IPv4Mask(0, 0, 255, 255)).String() == "0.0.0.2" {
		// Lowest IP address, assign supernode
		go listen()
	//BUCKET TESTING CODE
	b = newBucket()

	//Use line to find IP address for base node
	//fmt.Println(GetOutboundIP().String())
	if GetOutboundIP().String() == "172.19.0.2" {
		listen()

	} else {
		go sendLoop()
	}

	for {
		//fmt.Println("Alive") // Debug printout to ensure node is alive
		time.Sleep(time.Second / 2)
	}
}

func sendLoop() {
	networkPrefix1, _ := strconv.Atoi(strings.Split(netInfo.localIPAddr.String(), ".")[0])
	networkPrefix2, _ := strconv.Atoi(strings.Split(netInfo.localIPAddr.String(), ".")[1])
	supernodeAddr := net.IPv4(byte(networkPrefix1), byte(networkPrefix2), 0, 2)
	for {
		// Forever ping the supernode
		sendPing(supernodeAddr)

		buffer := bytes.NewBuffer(inputBytes[:length])
		decoder := gob.NewDecoder(buffer)
		err = decoder.Decode(&message)
		if err != nil {
			fmt.Println(err)
		}
		b.addToBucket(message)
		//NewNode(message.ID, message.IP)
		fmt.Println("Received message from ", senderAddr, "\n Packet IP: ", message.IP.String(), "\n Sender ID: ", message.ID)
	}
}

func send() {
	dest_addr := "172.19.0.2"
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
	sendPack.ID = node.ID
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
