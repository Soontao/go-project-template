package main

import (
	"log"

	"github.com/urfave/cli"
)

var commandEntry = cli.Command{
	Name:   "entry",
	Usage:  "program entry",
	Action: entry,
}

func entry(c *cli.Context) error {
	log.Println("hello golang.")
	return nil
}
