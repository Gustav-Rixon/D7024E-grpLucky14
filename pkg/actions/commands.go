package actions

import (
	"io"
	"net"
	"os"
	"os/exec"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

func Commands(msg *[]byte) {
	app := &cli.App{
		Name:  "cli",
		Usage: "do cli stuffff",
	}

	app.Commands = []cli.Command{
		{
			Name:  "exit",
			Usage: "Terminates the node",
			Action: func(c *cli.Context) error {
				terminate()
				return nil
			},
		},
		{
			Name:  "ping",
			Usage: "ping a node",
			Action: func(c *cli.Context) error {
				sendMessage(msg)
				return nil
			},
		},
		{
			Name:  "put",
			Usage: "Takes a single argument, the contents of the file you are uploading, and outputs the hash of the object, if it could be uploaded successfully.",
			Action: func(c *cli.Context) error {
				sendMessage(msg)
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Takes a hash as its only argument, and outputs the contents of the object and the node it was retrieved from, if it could be downloaded successfully",
			Action: func(c *cli.Context) error {
				sendMessage(msg)
				return nil
			},
		},
		{
			Name:  "add",
			Usage: "Add contact to Routing table",
			Action: func(c *cli.Context) error {
				sendMessage(msg)
				return nil
			},
		},
		{
			Name:  "getTable",
			Usage: "Get contact of Routing table",
			Action: func(c *cli.Context) error {
				sendMessage(msg)
				return nil
			},
		},
		{
			Name:  "findNode",
			Usage: "Does a ns lookup on node ID",
			Action: func(c *cli.Context) error {
				sendMessage(msg)
				return nil
			},
		},
		{
			Name:  "getid",
			Usage: "get node ID",
			Action: func(c *cli.Context) error {
				sendMessage(msg)
				return nil
			},
		},
		{
			Name:  "join",
			Usage: "join",
			Action: func(c *cli.Context) error {
				sendMessage(msg)
				return nil
			},
		},
		{
			Name:  "forget",
			Usage: "forgets object with key",
			Action: func(c *cli.Context) error {
				sendMessage(msg)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Info().Msgf("CRASH")
	}
}

// Executes kill -9 7. NOTE this is HARD codeded and expects the docker continer to have the PID 7
func terminate() {
	exec.Command("kill", "-9", "7").Run()
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
