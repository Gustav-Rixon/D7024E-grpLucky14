package main

import (
	"io"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

func main() {

	if len(os.Args) > 1 { // If a command was specified
		msg := StrArrayToByteArray(os.Args[1:])
		sendMessage(&msg)
	} else {
		//TODO: Print usage
		log.Print("Usage: To be done...")
	}
}

func StrArrayToByteArray(strs []string) []byte {
	return []byte(strings.Join(strs, " "))
}

func reader(wg *sync.WaitGroup, r io.Reader) {
	defer wg.Done()

	//TODO: Don't hardcode buffer size to 1024 bytes
	buf := make([]byte, 10000)
	n, err := r.Read(buf[:])
	if err != nil {
		return
	}
	log.Info().Msgf("Received response: %s", string(buf[:n]))
}

func sendMessage(msg *[]byte) {
	c, err := net.Dial("unix", "/tmp/echo.sock")
	defer c.Close()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Make sure reader is set up before writing
	go reader(&wg, c)

	_, err = c.Write(*msg)
	if err != nil {
		log.Error().Msgf("Failed to write to socket: %s", err.Error())
	}
	wg.Wait()
}
