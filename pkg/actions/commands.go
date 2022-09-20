package actions

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func Commands() {
	app := &cli.App{
		Name:  "hello",
		Usage: "say hello",
		Action: func(c *cli.Context) error {
			fmt.Println("hello")
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
