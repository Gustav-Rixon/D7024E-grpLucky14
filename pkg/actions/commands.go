package actions

import (
	"fmt"
	"log"
	"os"

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
				fmt.Println("get")
				return nil
			},
		},
		{
			Name:  "exit",
			Usage: "do the thing",
			Action: func(c *cli.Context) error {
				fmt.Println("exit")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
