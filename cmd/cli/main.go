package main

import (
	"log"
	"os"

	commands "github.com/AdhityaRamadhanus/gopatrol/commands/cli"
	"github.com/urfave/cli"
)

func cmdNotFound(c *cli.Context, command string) {
	log.Printf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		os.Args[0],
	)
	os.Exit(1)
}

func main() {
	// Init Config
	app := cli.NewApp()
	app.Name = "gopatrol-cli"
	app.Author = "Adhitya Ramadhanus"
	app.Email = "adhitya.ramadhanus@gmail.com"

	app.Commands = commands.Commands
	app.CommandNotFound = cmdNotFound // Inspired by docker machine
	app.Usage = "gopatrol cli "
	app.Version = "1.0.0"

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
