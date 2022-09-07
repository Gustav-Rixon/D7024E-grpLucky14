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

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func main() {

    //node := NewNode()

    s2 := rand.NewSource(int64(ip2int(GetOutboundIP())))
    r1 := rand.New(s2)
    //r := rand.Intn(256)
    //r2 := rand.Intn(256)

    for {
        r := r1.Intn(256)
        time.Sleep(2 * time.Second)
        fmt.Println(r)
    }
}