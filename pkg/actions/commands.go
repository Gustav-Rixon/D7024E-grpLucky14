package actions

import (
	"fmt"
	. "kademlia/internal/datastorage"
	. "kademlia/internal/node"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func Commands() {
	app := &cli.App{
		Name:  "cli",
		Usage: "do cli stuffff",
	}

	app.Commands = []cli.Command{
		{
			Name:  "put",
			Usage: "Takes a singel argument, the contents of the file you are uploading, and outputs the hash of the object, if it could be uploaded successfully.",
			Action: func(c *cli.Context) error {
				fmt.Println("Pls input file <This is just a string>:")
				var file string
				fmt.Scanln(&file)
				Insert(file)
				fmt.Println(GetKey(file))
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Takes a hash as its only argument, and outputs the contents of the object and the node it was retrieved from, if it could be downloaded successfully",
			Action: func(c *cli.Context) error {
				fmt.Println("Pls input key <hash>:")
				var key string
				fmt.Scanln(&key)
				fmt.Println(GetKey(key))
				return nil
			},
		},
		{
			Name:  "exit",
			Usage: "Terminates the node",
			Action: func(c *cli.Context) error {
				terminate()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// Executes kill -9 7. NOTE this is HARD codeded and expects the docker continer to have the PID 7
func terminate() {
	exec.Command("kill", "-9", "7").Run()
}
