package main

import (
    "fmt"
    "math/rand"
    "time"
    "log"
    "net"
    "encoding/binary"
)

type Node struct{
    ID  *KademliaID
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

//converts ip address to int
func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

//Returns random number, used in Kademlia ID generation
func getRandNum() int{    
    r := rGen.Intn(256)
    return r
}

//Random number generator, use to get random numbers between nodes
var rGen *rand.Rand

func main() {
    //initialize randomization of ID
    randSource := rand.NewSource(time.Now().UnixNano())
    rGen = rand.New(randSource)
    NodeID := NewRandomKademliaID()


    for {
        time.Sleep(2 * time.Second)
        
        fmt.Println(NodeID)
        
    }
}