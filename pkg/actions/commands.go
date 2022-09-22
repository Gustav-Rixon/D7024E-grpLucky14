package actions

import (
	"fmt"
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
			Usage: "do the thing",
			Action: func(c *cli.Context) error {
				fmt.Println("put")
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "do the thing",
			Action: func(c *cli.Context) error {
				fmt.Println("bajsa")
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
